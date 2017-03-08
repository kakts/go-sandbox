package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range はスライスやマップを反復処理するときに使う
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
