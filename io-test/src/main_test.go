package main

import (
	"bytes"
	"errors"
	"testing"
	"testing/fstest"
	"testing/iotest"
)

func TestCount(t *testing.T) {
	t.Run("File does not exist.", func(t *testing.T) {
		fs := fstest.MapFS{}
		want := 0

		got, _ := countFiles(fs)
		assertCount(t, got, want)
	})

	t.Run("Only one file exists.", func(t *testing.T) {
		fs := fstest.MapFS{
			"a.txt": {Data: []byte("Hello, world!")},
		}
		want := 1
		got, _ := countFiles(fs)
		assertCount(t, got, want)
	})

	t.Run("Two files exist.", func(t *testing.T) {
		fs := fstest.MapFS{
			"a.txt": {Data: []byte("Hello, world!")},
			"b.txt": {Data: []byte("Hello, world!")},
		}
		want := 2
		got, _ := countFiles(fs)
		assertCount(t, got, want)
	})
}

func assertCount(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestToUpperWithErrReader(t *testing.T) {
	// 読み込み処理が失敗するio.Reader
	src := iotest.ErrReader(errors.New("failed read file."))
	dst := bytes.NewBufferString("")

	err := toUpper(src, dst)
	if err == nil {
		t.Errorf("no error occurred")
	}
}

func TestToUpperWithTimeoutReader(t *testing.T) {
	// タイムアウトを発生させるio.Reader
	src := iotest.NewReadLogger("read log", iotest.TimeoutReader(bytes.NewReader([]byte("hello"))))

	dst := bytes.NewBufferString("")

	err := toUpper(src, dst)

	if err == nil {
		t.Errorf("no error occurred")
	}
}
