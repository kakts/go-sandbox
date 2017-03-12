package main

import "fmt"

func main() {
	// 空のインターフェース
	// 任意の方の値を保持できる
	// 空のインターフェースは、未知の型の値を扱うコードで使用できる
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)", i, i)
}
