package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		// handle Error
		fmt.Println("err")
	}
	fmt.Printf(resp.Status)
	if resp.StatusCode == 200 {
		fmt.Println("success")
	}

	// The http client should be closed finally.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
