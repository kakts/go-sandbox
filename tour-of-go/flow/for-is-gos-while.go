package main

import "fmt"

func main() {
  sum := 1
  // セミコロン省略可能
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
