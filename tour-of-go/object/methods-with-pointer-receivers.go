package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// ポインタレシーバを使う理由
// 1: メソッドがレシーバが指す先の変数を変更する
// 2: メソッドの呼び出し毎に変数のコピーを避ける
//    レシーバが大きな構造体である場合に非常に効率がよくなる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scalling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scalling: %+v, Abs: %v\n", v, v.Abs())
}
