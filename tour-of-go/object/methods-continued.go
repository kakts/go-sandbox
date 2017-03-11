package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// メソッドの宣言は、レシーバ型(MyFloat)が同じパッケージにある必要がある
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
