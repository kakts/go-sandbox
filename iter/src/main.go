package main

import "fmt"

// iterator はイテレータ関数
// yieldを呼び出すたびにイテレータ関数の実行が一時停止され、
// 呼び出し元のforループに値が渡される
func iterator(yield func(int) bool) {
	fmt.Println("iterator yield is calling yield(1)")
	yield(1)
	fmt.Println("iterator yield is calling yield(2)")
	yield(2)
	fmt.Println("iterator yield is calling yield(3)")
	yield(3)
}



func main() {
	for i := range iterator {
		fmt.Printf("for loop executed: %i ", i)
		fmt.Println(i)
	}
}