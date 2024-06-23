package slice

import (
	"testing"
)

// initで100万個の要素にランダムに3文字を含むstringのスライスを作成
var foos = make([]string, 1000000)

/**
 * スライスの初期化におけるベンチマーク比較
 * makeによるスライス初期化でlen capの指定の有無によってパフォーマンス
 * がどのように変わるかを確認する
 */

// makeでスライスを作成時に、lenのみ0指定を行った場合
func Benchmark_sliceMakeOnlyWithSize(b *testing.B) {
	// sliceMakeOnlyWithSizeに渡すために100万個の要素を持つstringのスライスを作成
	sliceMakeOnlyWithSize(foos)
}

func Benchmark_sliceMakeWith(b *testing.B) {
	// sliceMakeWithSettingSizeAndCapに渡すために100万個の要素を持つstringのスライスを作成
	sliceMakeWithSettingSizeAndCap(foos)
}

func Benchmark_sliceMakeOnlyWithSettingSizeFoosLen(b *testing.B) {
	// sliceMakeOnlyWithSettingSizeFoosLenに渡すために100万個の要素を持つstringのスライスを作成
	sliceMakeOnlyWithSettingSizeFoosLen(foos)
}
