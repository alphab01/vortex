package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	get, _ := http.Get("http://127.0.0.1:4000")
	fmt.Println(get.Body)
	s, _ := io.ReadAll(get.Body)
	fmt.Println(string(s))
}
