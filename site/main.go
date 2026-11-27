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
	s := experiment{Lib: "iostream", Date: "27.11.2026", Gpu: "Nvidia GeForce RTX 5090", Desc: "test", Gpucomm: "111", Libcomm: "222", Sim: "testsim", Res: "Good"}
	d := template.Must(template.New("main").Parse(`<div>
		<h1>SparceLinearAlgebra - Vortex</h1>
		<p>description</p>
		<hr>
		<p>our experiments:</p>
		<table>
			<tr>
				<p> {{ .Lib}} </p>
				<p> {{ .Date}} </p>
				<p> {{ .Gpu}} </p>
				<p> {{ .Desc}} </p>
				<p> {{ .Gpucomm}} </p>
				<p> {{ .Libcomm}} </p>
				<p> {{ .Sim}} </p>
				<p> {{ .Res}} </p>
			</tr>
			<p> bye </p>
		</table>
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
