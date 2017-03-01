package main

import "fmt"

// 関数宣言の行で戻り値の変数名もわかるので可読性が上がる
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x // XXX
  return
}

func main() {
  fmt.Println(split(17))
}
