package goroutine

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func SimpleChannel() {
	stringStream := make(chan string)

	go func() {
		stringStream <- "Hello channels!"
	}()

	// チャネルからデータを受信するまでブロック
	fmt.Println(<-stringStream)
}

func ChannelReceiveCheck() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	// チャネルの第2戻り値はチャネルが受信成功したかをチェックする
	// falseの場合 チャネルが閉じられていてこれ以上受信できないことを示す
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
}

func ChannelLoop() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// チャネルに対するrangeループは、チャネルが閉じられるまでデータを受信し続ける
	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

func ReleaseChannelForMultiGoroutine() {
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// チャネルからの読見込めるようになるまでブロック
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	// チャネルを閉じる　これにより全てのgoroutineが同時に開始される
	close(begin)
	wg.Wait()
}

func BufferedChannel() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	// バッファ付きチャネル 4つまではブロッキングなしでデータを送受信できる
	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "SEnding: %d\n", i)
			// バッファ付きチャネルへの送信 4つまではブロッキングなし
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received: %v\n", integer)
	}
}

func ChanOwner() {
	// 読み取り専用チャネルを初期化して返す
	chanOwner := func() <-chan int {
		// バッファ付きチャネルの初期化
		resultStream := make(chan int, 5)

		// resultStreamへの書き込みを行うgoroutineを起動
		// goroutineよりも先にresultStreamを生成した
		// 外側のchanOwner関数によってカプセル化している
		go func() {
			// resultStreamを使い終わったらクローズ
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		// 読み込み専用チャネルの返却
		return resultStream
	}

	resultStream := chanOwner()

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}
