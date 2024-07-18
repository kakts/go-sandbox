package goroutine

import (
	"fmt"
	"time"
)

func SimpleSelect() {
	start := time.Now()

	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	// channel cの読み込みを試みる クローズされると、このcaseが実行される
	case <-c:
		fmt.Printf("Unblocked %v later\n", time.Since(start))
	}
}
