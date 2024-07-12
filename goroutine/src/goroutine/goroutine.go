package goroutine

import (
	"fmt"
	"sync"
)

// SayHello は"Hello"と出力する
// sync.WaitGroupによって、goroutineの終了を待つ
func SayHelloWithWaiting() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}

	// ここで指定した数だけgoroutineの終了を待つ。
	// カウンターが0になったらwait状態になっているこの関数でのgoroutineが解放される
	wg.Add(1)
	go sayHello()
	wg.Wait()
}
