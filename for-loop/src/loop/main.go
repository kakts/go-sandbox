package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
		"d": 400,
		"e": 500,
	}

	for k, v := range m {
		fmt.Printf("%s : %d \n", k, v)
	}
}
