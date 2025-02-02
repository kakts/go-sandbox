package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

// countFiles は ファイル数をカウントする
func countFiles(fsys fs.FS) (int, error) {
	files, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	f, err := os.Create("test.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name(), "is a directory")
		}
		fmt.Println(file.Name())
	}

	return len(files), nil
}

// toUpper は src(io.Reader)から読み込んだ文字列を大文字に変換した上で、dst(io.Writer)に書き込む
func toUpper(src io.Reader, dst io.Writer) error {
	b, err := io.ReadAll(src)
	if err != nil {
		return err
	}
	upperStr := strings.ToUpper(string(b))

	// dst writer に upperStr を書き込む
	_, err = io.WriteString(dst, upperStr)

	return err
}

func main() {

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// 書き込み先
	dst, err := os.Create("dst.txt")
	if err != nil {
		panic(err)
	}

	err = toUpper(f, dst)
	if err != nil {
		panic(err)
	}
}
