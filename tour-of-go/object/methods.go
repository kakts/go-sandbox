package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// func と メソッド名(Abs)の間にメソッドを設定したい型の引数を指定してメソッドを表現する
// メソッド ＝ レシーバ引数を伴う関数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
