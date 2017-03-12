package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 型アサーション
	// この文はインターフェースiが具体的な型(string)を保持し、元になるstringの値をsに代入することを主張する
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	f = i.(float64)
	fmt.Println(f)
}
