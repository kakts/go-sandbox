package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Cond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]any, 0, 10)

	// 指定した秒数後にqueueから要素を取り出す
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()

		// 条件を待っているgoroutineに何かが起きたことを知らせる
		// この場合、queueに要素が追加されたことを知らせる
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		// クリティカルセクション
		c.L.Lock()
		for len(queue) == 2 {
			// 呼び出したgoroutineで条件が満たされるまで待つ
			// c.Signal()によって条件が満たされると、この関数は返る
			c.Wait()
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func CondButton() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			fmt.Println("Waiting for broadcast")
			c.Wait()
			fmt.Println("Broadcast signal emitted")
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// BroadCastにより、c.Wait()で待っているgoroutine全てに対して通知を行う
	fmt.Println("Before Broadcast")
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
