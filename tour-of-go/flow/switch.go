package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on")

  // osで切り替え
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X")
    // breakを書句必要がない
    // breakせずに通したい場合はfallthroughを書く
  case "linux":
    fmt.Println("Linux")
  default:
    fmt.Println("%s.", os)
  }
}
