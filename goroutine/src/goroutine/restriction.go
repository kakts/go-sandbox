package goroutine

import (
	"bytes"
	"fmt"
	"sync"
)

// AdhocRestriction アドホック拘束の例
func AdhocRestriction() {

	/**
	 * dataは規約によりloopData関数内からのみアクセスされている
	 * しかし改修により、dataに対して他の関数からアクセスする可能性がでてきてしまうのが課題
	 */
	data := make([]int, 4)

	// 読み取り専用のチャネルを受け取りそのチャネルに対してデータを送信する
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func LexicalRestriction() {
	/**
	 * レキシカル拘束の例
	 * レキシカルスコープを使って適切なデータと並行処理のプリミティブ
	 * だけを複数の並行プロセスが使えるように公開する
	 */

	// 送信専用のチャネルを返す関数
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// 送信専用のチャネルを受け取り、そのチャネルからデータを受信する関数
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	// chanOwner関数を呼び出し、その結果をconsumer関数に渡す
	results := chanOwner()
	consumer(results)
}

// BytesRestriction 並行安全でないbytes.Bufferを使った拘束の例
func BytesRestriction() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")
	// dataを2つに分割し、goroutineに渡す
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
}
