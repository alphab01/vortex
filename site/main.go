package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type test struct {
	S1 string
	S2 string
}

func h(w http.ResponseWriter, r *http.Request) {
	s := test{S1: "Hello", S2: "World"}
	d := template.Must(template.New("main").Parse(`<div>
		<h1>{{ .S1}}</h1>
		<p>{{ .S2}}</p>
</div>`))
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
