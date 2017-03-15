package main

import "fmt"

// channel
// ch <- v    vをチャネルchへ送信する
// v := <-ch  chから受信した変数をvへ割り当てる

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// intを受け取るchannelを作成する
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

}
