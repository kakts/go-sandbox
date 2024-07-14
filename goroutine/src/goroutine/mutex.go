package goroutine

import (
	"fmt"
	"sync"
)

// MutexLock はsync.Mutexを使って排他制御を行う
// sync.WaitGroupによって全てのgoroutineの終了を待つ
// countをインクリメント・デクリメントするが、それぞれsync.Mutexによる排他制御を行っている
func MutexLock() {
	var count int
	var lock sync.Mutex

	// countをインクリメントする関数
	// クロージャによってlockを使って排他制御を行っている
	increment := func() {
		lock.Lock()
		// Unlockはdeferで行うことにより、途中でpanicにあったとしても確実に呼び出される
		// 呼び出しが失敗するとデッドロックが発生してしまう。
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	// インクリメントする
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// デクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()

	fmt.Printf("Arithmetic complete. ")
}
