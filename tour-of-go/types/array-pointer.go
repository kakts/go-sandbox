package main

import "fmt"

func main() {
  arr := [4]int{1, 2, 3, 4}
  argVal(arr)
  fmt.Println(arr)

  argPointer(&arr)
  fmt.Println(arr)
}

// 引数は値渡しになる
// 引数の値をコピーしたものを関数内で使っている
func argVal(arr [4]int) {
  fmt.Println("---argVal---")
  fmt.Println(arr)
  arr[2] = 14
  fmt.Println(arr)
  fmt.Println("---argVal finish---")
}

// 引数はポインターを渡しているため、参照渡しになる
// 関数内で値を変えると、呼び出し元でも変更がかかる
func argPointer(arr *[4]int) {
  fmt.Println("---argPointer---")
  fmt.Println(arr)
  arr[2] = 14
  fmt.Println(arr)
  fmt.Println("---argPointer finish---")
}
