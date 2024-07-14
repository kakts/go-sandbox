package goroutine

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
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

// RWMutexLock はsync.RWMutexとsync.Mutexの実行時間を比較する
func RWMutexLock() {
	// 第2引数の sync.Lockerは Lock()とUnlock()を持つインターフェース
	// MutexとRWMutexはこのインターフェースを満たしている
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()

		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		// すべてのgoroutineが終了するまで待つ
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}
