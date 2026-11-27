package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type experiment struct {
	Lib string
	Date string
	Gpu string
	Desc string
	Gpucomm string
	Libcomm string
	Sim string
	Res string
}

type test struct {
	S1 string
	S2 string
}

func h(w http.ResponseWriter, r *http.Request) {
	s := test{S1: "Hello", S2: "World"}
	d := template.Must(template.New("main").Parse(`<div>
		<h1>SparceLinearAlgebra - Vortex</h1>
		<p>description</p>
		<hr>
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
