package main

import "fmt"

func main() {
	// len=10の配列スライスの初期化
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// 使わない値は_に代入することで捨てることができる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
