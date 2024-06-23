package slice

// sliceMakeOnlyWithSize は、makeでスライスを作成時に、lenのみ0で指定を行った場合
func sliceMakeOnlyWithSize(foos []string) []string {

	bars := make([]string, 0)

	for _, foo := range foos {
		bars = append(bars, foo)
	}

	return bars
}

// sliceMakeWithSettingSizeAndCap は、makeでスライスを作成時に、len capの指定を行った場合
func sliceMakeWithSettingSizeAndCap(foos []string) []string {
	n := len(foos)
	// スライスの長さを0、capをfoosの要素数に設定
	bars := make([]string, 0, n)

	for _, foo := range foos {
		bars = append(bars, foo)
	}

	return bars
}

// sliceMakeOnlyWithSettingSizeFoosLen は、makeでスライスを作成時に、lenのみの指定を行った場合
func sliceMakeOnlyWithSettingSizeFoosLen(foos []string) []string {
	n := len(foos)
	// スライスの長さをfoosの要素数を指定 capの指定なし
	bars := make([]string, n)

	for i, foo := range foos {
		// スライスの初期化時にsizeをfoosの長さで初期化しているため
		// すでにn個の要素の領域が確保されていて、ゼロ値に初期化されている
		// そのため、appendでなく、bars[i]に値を代入する
		bars[i] = foo
	}

	return bars
}
