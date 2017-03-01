package main

import "fmt"

func main() {
  var i = 1
  defer fmt.Println("world. %d", i)
  i++
  defer fmt.Println("world2. %d", i)
  fmt.Println("hello %d", i)
}
