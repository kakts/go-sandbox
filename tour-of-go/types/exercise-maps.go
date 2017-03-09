package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := map[string]int{}
	list := strings.Fields(s)
	fmt.Println(list)
	for i := range list {
		str := list[i]
		_, ok := countMap[str]
		if !ok {
			countMap[str] = 1
		} else {
			countMap[str] += 1
		}

	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
