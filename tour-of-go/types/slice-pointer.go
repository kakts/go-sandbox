package main

import (
  "fmt"
)

func main() {
  arr := [4]int{1, 2, 3, 4}
  fmt.Println(arr)
  sl := []int{1, 2, 3, 4}
fmt.Println(sl) // [1 2 100 4]と表示
  sl = sl[1:4]
  fmt.Println("len, %d cap %d", len(sl), cap(sl))

  // 別の変数に参照を渡す
  sl2 := sl
  // 0番目の要素に10を代入する sl2もslと同じ参照なので、sl2[0]も10になる
  sl[0] = 10
  fmt.Println("sl: len, %d cap %d", len(sl), cap(sl), sl)
  fmt.Println("sl2: len, %d cap %d", len(sl2), cap(sl2), sl2)

  // appendする
  // 既にcapが4, lenが4の状態でappendするので、内部で新たにcapを２倍にした文のメモリをallocし、返却している
  // なので、sl2とは指しているポインターのアドレスが異なっている
  sl = append(sl, 20)
  sl[0] = 100
  fmt.Println("sl: len, %d cap %d", len(sl), cap(sl), sl)
  fmt.Println("sl2: len, %d cap %d", len(sl2), cap(sl2), sl2)

sl3 := []int{1, 2, 3, 4}
fmt.Println("sl: len, %d cap %d", len(sl3), cap(sl3), sl3)
sl3 = append(sl3, 5)
fmt.Printf("sl: len, %d cap %d", len(sl3), cap(sl3), sl3)
}
