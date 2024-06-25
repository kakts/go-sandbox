package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("src/template/login.gtpl")
		if err != nil {
			log.Fatal("ParseFiles: ", err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login)

	// HTTPSサーバーを起動
	err := http.ListenAndServeTLS(":8080", "./server.crt", "./server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
