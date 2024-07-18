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

func MultiSelect() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\n", c1Count)
	fmt.Printf("c2Count: %d\n", c2Count)
}
