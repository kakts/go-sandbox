package main

import "fmt"

func main() {
	// value:intのマップを作る
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかは2つ目の値で確認できる
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
