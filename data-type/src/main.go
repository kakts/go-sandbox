package main

import "fmt"

func main() {

	fmt.Println(octaSum())
}

// 10進数と8進数の和を求める
// 0から始まる整数リテラルは８進数として扱われるため、意図しない計算結果となる
func octaSum() int {
	sum := 100 + 010
	return sum
}
