package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
	s := "html template"
	d, _ := template.New("main").Parse("<h1>{{ .}}</h1>")
	d.Execute(w, s)
}

func main() {
	serv := http.NewServeMux()
	serv.HandleFunc("/", h)
	err := http.ListenAndServe(":4000", serv)
	if (err != nil) {
		fmt.Println(err)
	}
}
