package goroutine

import "fmt"

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
