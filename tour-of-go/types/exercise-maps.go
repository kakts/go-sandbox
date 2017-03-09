package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  countMap := map[string]int{}
  list := strings.fields(s)
  for data := range list {
    count, ok := map[data]
    if !ok {
      map[data] = 1
    } else {
      map[data] += 1
    }

  }
  return countMap
}

func main() {
  wc.Test(WordCount)
}
