package loop

import "fmt"

// MapLoop マップをループさせる場合 順序は保証されないことを確認する
// see: https://go.dev/blog/maps#iteration-order
func MapLoop() {
	// Mapのループ
	m := map[string]int{
		"apple":  100,
		"banana": 200,
		"cherry": 300,
	}

	for k, _ := range m {
		fmt.Print(k)
	}

}
