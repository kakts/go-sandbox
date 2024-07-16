package goroutine

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func Pool() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance.")
			return struct{}{}
		},
	}

	// sync.Poolからインスタンスを取得
	// なければNewメンバ変数で指定した関数を呼び出してインスタンスを生成
	myPool.Get()
	instance := myPool.Get()

	myPool.Put(instance)
	myPool.Get()
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("Cannot listen: %v", err)
		}

		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}

			connectToService()
			fmt.Println(conn, "")
			conn.Close()
		}
	}()

	return &wg
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("Cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("Cannot read: %v", err)
		}

		conn.Close()
	}
}
