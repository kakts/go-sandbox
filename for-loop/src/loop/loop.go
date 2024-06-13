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

// AddDataToMapInLoop ループ中にマップのエントリが追加される場合、ループ中に取り出される可能性があったり、スキップされたりする
// どういう処理になるかは実行ごとに異なる
// see: https://go.dev/ref/spec#For_statements
func AddDataToMapInLoop() {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		fmt.Println(k, v)
		if v {
			m[10+k] = true
		}
	}
	fmt.Println(m)
}
