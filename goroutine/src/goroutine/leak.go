package goroutine

import (
	"fmt"
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
