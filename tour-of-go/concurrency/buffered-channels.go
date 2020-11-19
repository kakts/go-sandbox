package main

import "fmt"

func main() {
	/**
	 * チャネルはバッファとしても使える
	 * バッファを持つチャネルを初期化するにはmakeの2つ目の引数にバッファの長さを与える
	 */
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 これを行うとデッドロックする
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
