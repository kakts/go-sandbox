package goroutine

import (
	"fmt"
	"sync"
	"time"
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
	wg.Add(3)
	count := 0
	for {
		if count == 3 {
			break
		}
		go sayHello()
		time.Sleep(1 * time.Second)
		count++
	}
	// go sayHello()
	wg.Wait()
}

// SayHelloWithClosure クロージャを使って変数をgoroutine内で変更するパターン
func SayHelloWithClosure() {
	fmt.Println("=========SayHelloWithClosure()=========")
	var wg sync.WaitGroup
	salutation := "Hello"
	wg.Add(1)

	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

// SayHelloWithClosure2 クロージャを使ってfor ループ変数をgoroutine内で変更するパターン
// go v1.21まではループの最後の要素が出力される
func SayHelloWithClosure2() {
	var wg sync.WaitGroup
	fmt.Println("=========SayHelloWithClosure()=========")
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)

		// saluationはgoroutine外で定義され、スコープ害となっている
		// ここで関数の引数にsalutationを渡せば文字列がコピーされる
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
			fmt.Printf("&salutation: %p\n", &salutation)
		}()
	}
	wg.Wait()
}
