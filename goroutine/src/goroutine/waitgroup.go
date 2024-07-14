package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping..")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping..")
		time.Sleep(2)
	}()

	// wait until the all goroutines are finished
	wg.Wait()
	fmt.Println("All goroutines are finished")
}

func WaigGroupMulti() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup

	// goroutine呼び出し前に1度で全てのgoroutineの数を指定する
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
