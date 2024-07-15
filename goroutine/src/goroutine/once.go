package goroutine

import (
	"fmt"
	"sync"
)

func Once() {
	var count int

	increment := func(index int) {
		fmt.Printf("increment is called. index: %d\n", index)
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup

	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer increments.Done()
			once.Do(func() {
				increment(i)
			})
		}(i)
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
