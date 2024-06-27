package main

import (
	"net/http"

	"github.com/kakts/go-sandbox/http/src/handler"
)

func main() {

	http.HandleFunc("GET /comments", handler.GetComments)
	http.HandleFunc("POST /comment", handler.PostComment)

	http.ListenAndServe(":8080", nil)
}
