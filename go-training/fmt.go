package main

import (
	"fmt"

	"rsc.io/quote"
)

// T sample type
type T struct {
	name  string // a
	value int    // b
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	var a = T{name: "a", value: 1}
	fmt.Println(a.name, a.value)
	fmt.Println(sum(10, 5))
	fmt.Println(quote.Go())
}
