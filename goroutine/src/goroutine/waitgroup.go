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
