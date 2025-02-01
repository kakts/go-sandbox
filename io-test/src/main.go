package main

import (
	"fmt"
	"io/fs"
	"os"
)

// countFiles は ファイル数をカウントする
func countFiles(fsys fs.FS) (int, error) {
	files, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name(), "is a directory")
		}
		fmt.Println(file.Name())
	}

	return len(files), nil
}

func main() {
	count, err := countFiles(os.DirFS("src"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}
