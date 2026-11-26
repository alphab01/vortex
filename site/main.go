package main

import (
	"fmt"
	"net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Vortex"))
}

func main() {
	serv := http.NewServeMux()
	serv.HandleFunc("/", h)
	err := http.ListenAndServe(":4000", serv)
	if (err != nil) {
		fmt.Println(err)
	}
}
