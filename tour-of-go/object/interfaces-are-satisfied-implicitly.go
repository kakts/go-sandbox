package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S  string
	S2 string
}

// このメソッドは T が interface Iを実装していることを表している
// 明示的にimplementsを宣言しなくて良い
func (t T) M() {
	fmt.Println(t.S, t.S2)
}

func main() {
	var i I = T{"hello", "world"}
	i.M()
}
