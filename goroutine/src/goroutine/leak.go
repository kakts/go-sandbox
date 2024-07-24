package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

func GoroutineLeak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// nilチャネルを渡すことで、doWork内部で絶対に文字列が書き込まれることがない
	// そのため doWorkを含むgoroutineはプロセスが生成する限りずっとメモリ内に残る
	doWork(nil)

	fmt.Println("Done.")
}

func GoroutineLeakFix() {

	/**
	 * goroutine親子間でキャンセルのシグナルをやり取りできるようにする
	 */
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		fmt.Println("Calling Goroutine in doWork.")
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done: // doWorkの呼び出し側でdoneチャネルが閉じられると、このcaseが実行される
					fmt.Println("done is closed.")
					return
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})

	// 第2引数 読み取り専用チャネルをnilにすることで、内部処理で書き込みが行われないようにする
	terminated := doWork(done, nil)

	go func() {
		// 1秒後にキャンセルシグナルを送信
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	// terminatedチャネルが閉じられるまで待機
	<-terminated
	fmt.Println("Done.")
}

// goroutineがチャネルに対して書き込みを行おうとしてブロックする例
func WriteChannelLeak() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			// ここのfmt.Printlnは決して実行されない
			// 利用側で3回の読み込みしてそのあとは読み込まないため
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				// 3回のみ送信できるが、その後は呼び出し側で読み込まないため、ブロックされる
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		// ここで3回チャネルから読み込んだ後に、goroutine側に停止していいと伝える方法がない
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

// WriteChannelLeakの問題を解決するための修正
func WriteChannelLeakFix() {
	/**
	 * 生産者のgoroutineに終了を伝えるチャネル doneを提供することで、
	 * 呼び出し元のgoroutineからチャネルからの読み込み停止を通知する
	 */
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int(): // 何もしない
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	fmt.Println("Cancelling the randStream goroutine...")
	// 3回読み込んだ後にdoneチャネルを閉じることで、goroutineに終了を通知
	close(done)

	// 処理が実行中であることをシミュレート
	time.Sleep(1 * time.Second)
}
