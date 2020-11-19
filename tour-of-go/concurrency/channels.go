package main

import "fmt"


/**
 * channel
 * チャネルオペレータの<- を用いて値の送受信ができる通り道
 * ch <- v : vをチャネル ch へ送信する
 * v := <- ch : chから受信した変数を v へ割り当てる
 * 
 * データは矢印の方向に流れる
 */
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// チャネルは通常片方が準備できるまでは送受信はブロックされる
	// これにより明確なロックや条件変数がなくてもgoroutineの動機を可能にする
	// ここでは2つのgoroutine間で作業を分配する
	// 両方のgoroutineで計算完了すると最終結果が計算される
	go sum(s[:len(s) / 2], c)
	go sum(s[len(s) / 2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x + y)
}