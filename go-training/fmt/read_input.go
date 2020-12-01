package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFromStdinOneLine reads from stdin only one line and print it.
func ReadFromStdinOneLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please type something. : ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("input is", input)
	return input
}

func main() {
	var inputFromStdin = ReadFromStdinOneLine()
	fmt.Println("result from ReadFromStdinOneLine is ", inputFromStdin)
}
