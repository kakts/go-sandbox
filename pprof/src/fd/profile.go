// fd はpprof.Profileの機能を用いたファイルディスクリプタのプロファイリングを提供します。
package fd

import (
	"fmt"
	"os"
	"runtime/pprof"
)

// pprof.NewProfile は新しいプロファイルを作成します
// ここではグローバル変数として使用する
// プロファイル名は一意である必要がある
var fdProfile = pprof.NewProfile("fd.inuse")

// File はファイルディスクリプタのライフタイムを追跡するos.Fileのラッパー
type File struct {
	*os.File
}

// Open はファイルを開き、fdProfileで追跡します
func Open(name string) (*File, error) {
	// 既存のfdを記録するためにos.Openを利用
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file. err: %v", err)
	}

	// オプジェクトを追跡するためにfdProfileに追加
	// 第2引数はスタックトレースでスキップする呼び出し回数を指定
	// ここではOpen関数を参照することにするため、スタックフレームを2つスキップすることとする
	fdProfile.Add(f, 2)
	return &File{File: f}, nil
}

// Close はファイルを閉じ、プロファイルを更新します
func (f *File) Close() error {
	defer fdProfile.Remove(f.File)
	return f.File.Close()
}

// Write は今開いているファイルディスクリプタのプロファイルをpprof形式でファイルに保存する
func Write(profileOutPath string) error {
	out, err := os.Create(profileOutPath)
	if err != nil {
		return fmt.Errorf("Failed to create file. err: %v", err)
	}
	if err := fdProfile.WriteTo(out, 0); err != nil {
		_ = out.Close()
		return fmt.Errorf("Failed to write profile. err: %v", err)
	}
	return out.Close()
}
