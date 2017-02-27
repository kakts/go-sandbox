package main

import (
  "fmt"
  "reflect"
)
// compile error
l := 1
func main() {

  var i int = 1
  j := "1"
  k := 1
  fmt.Println(reflect.TypeOf(i))
  fmt.Println(reflect.TypeOf(j))
  fmt.Println(reflect.TypeOf(k))
}
