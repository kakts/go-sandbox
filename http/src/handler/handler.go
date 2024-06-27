package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	// mutex はコメントを保持するスライスの排他制御を行う
	// このパッケージ内で有効
	mutex *sync.RWMutex
	// comments はコメントを保持するスライス
	comments []Comment
)

func init() {
	// mutexを&sync.RWMutex{}で初期化
	mutex = &sync.RWMutex{}
	// commentsをmake([]Comment, 0, 100)で初期化
	comments = make([]Comment, 0, 100)
}

// Comment はコメントを表す構造体
type Comment struct {
	Message  string
	UserName string
}

// GetComments はコメントをjson形式で返す
func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 読み取り時に書き込みがあることを考慮し、ロックする
	mutex.RLock()
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	mutex.RUnlock()
}

// PostComment はコメントを受け取り、commentsに追加する
func PostComment(w http.ResponseWriter, r *http.Request) {
	var c Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
		return
	}

	// 同時に複数アクセスを防ぐためにロック
	mutex.Lock()
	comments = append(comments, c)
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"created"}`))
}
