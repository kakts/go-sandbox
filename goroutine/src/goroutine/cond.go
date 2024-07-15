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
