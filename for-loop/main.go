package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

// for rangeでのループ変数は、イテレーションごとに変数が再利用されるため、ポインタを使っても値が更新される
func main() {

	users := []*User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i, user := range users {
			fmt.Println("i = ", i)
			user.Age = 1000
			fmt.Printf("user: %v\n", *user)

			fmt.Printf("pointer of user: %p\n", &user)
		}
	}()

	wg.Wait()
}

func loopSliceOfPointer() {
	users := []*User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	var wg = sync.WaitGroup{}
	wg.Add(1)

	for i, user := range users {
		fmt.Println("i = ", i)
		fmt.Printf("user: %v\n", *user)

		// go1.21だと rangeでのループ変数のアドレスはイテレーションごとに変わらない
		// go1.22だと rangeでのループ変数のアドレスはイテレーションごとに変わる
		fmt.Printf("pointer of user: %p\n", &user)
	}
}
