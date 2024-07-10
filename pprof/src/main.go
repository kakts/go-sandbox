package main

import (
	"io"
	"log"
	"sync"

	"github.com/kakts/go-sandbox/pprof/src/fd"
)

/**
 * 112個のファイルを開くアプリをシミュレートする
 */
type TestApp struct {
	files []io.ReadCloser
}

// Close TestAppの全てのファイルを閉じる
func (a *TestApp) Close() {
	for i, cl := range a.files {
		// TODO err check
		_ = cl.Close()
		a.files[i] = nil
	}
	a.files = a.files[:0]
}

func (a *TestApp) open(name string) {
	// fd.Openを使ってファイルを開くと、副作用としてプロファイルに記録が開始される
	f, _ := fd.Open(name)
	a.files = append(a.files, f)
}

func (a *TestApp) OpenSingleFile(name string) {
	a.open(name)
}

func (a *TestApp) OpenTenFiles(name string) {
	for i := 0; i < 10; i++ {
		a.open(name)
	}
}

func (a *TestApp) Open100FilesConcurrently(name string) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// goroutineで10個ずつファイルを開く
		go func() {
			a.OpenTenFiles(name)
			wg.Done()
		}()
	}
	wg.Wait()
}

// pprofのプロファイルのテスト用のダミーファイル
const targetFilePath = "/dev/null"

func main() {
	a := &TestApp{}
	defer a.Close()

	// プロファイリングの動作を確認するために、十個のファイルを開いては閉じるを10回繰り返す
	for i := 0; i < 10; i++ {
		a.OpenTenFiles(targetFilePath)
		a.Close()
	}

	// 最後のCloseの後は、これより以下のファイルしかプロファイルに使われない
	f, _ := fd.Open(targetFilePath)
	a.files = append(a.files, f)

	a.OpenSingleFile(targetFilePath)
	a.OpenTenFiles(targetFilePath)
	a.Open100FilesConcurrently(targetFilePath)

	if err := fd.Write("fd.pprof"); err != nil {
		log.Fatal(err)
	}
}
