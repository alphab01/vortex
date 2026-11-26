package main

import (
	"fmt"
	"net/http"
)

func main() {
	get, _ := http.Get("http://127.0.0.1:4000/")
	fmt.Println(get)
}
