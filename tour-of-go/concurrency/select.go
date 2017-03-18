package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// selectステートメントは、goroutineを複数の通信操作で待たせる
		// 複数あるcaseのいずれかが準備できるようになるまでブロックし、
		// 準備ができたcaseを実行する
		// 複数のcaseが準備できている場合 ランダムに選択される
		select {
		case c <- x:
			fmt.Println("c")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
