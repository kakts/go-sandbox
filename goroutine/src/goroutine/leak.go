package goroutine

import "fmt"

func GoroutineLeak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()

		return completed
	}

	// nilチャネルを渡す
	doWork(nil)

	fmt.Println("Done.")
}
