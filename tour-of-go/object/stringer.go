package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	// a と zのインスタンスを作る
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	// Person structはString()を持っているため、stringとして表現することができる
	// printするときに、引数に渡したオブジェクトがこのstringerインターフェースを持っているかを確認している
	fmt.Println(a, z)
}
