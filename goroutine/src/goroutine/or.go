package goroutine

func OrChannel() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	// チャネルの可変長引数のスライスを受け取り、1つのチャネルを返す
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			// 引数がない場合はnilを返す
			return nil
		case 1:
			// 引数が1つの場合はそのまま返す
			return channels[0]
		}

		orDone := make(chan interface{})
		// goroutineにより、ブロックすることなく作ったチャネルにメッセージを受け取れるようにする
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				// orへの再起呼び出しは少なくとも2つのチャネルを持つ
				// goroutineの数を制限するために2ツッしかチャネルがなかった場合の特別な条件を設定する
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				// スライスの3番目以降のチャネルから再起的にorチャネルを作成して、そこからselectを実行する
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}
