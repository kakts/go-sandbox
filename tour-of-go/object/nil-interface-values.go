package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)

	// Mを実装していないため、ランタイムエラーになる
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}
