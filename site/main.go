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

type pair struct {
	AA []test
	BB experiment
}

func h(w http.ResponseWriter, r *http.Request) {
	//s := experiment{Lib: "iostream", Date: "27.11.2026", Gpu: "Nvidia GeForce RTX 5090", Desc: "test", Gpucomm: "111", Libcomm: "222", Sim: "testsim", Res: "Good"}
	var s []test
	s = append(s, test{S1: "s1", S2: "s2"})
	s = append(s, test{S1: "s3", S2: "s4"})
	d := template.Must(template.New("main").Parse(`<head>
			<style>
				th, td {
				border: 1px solid #111111}
			</style>
		</head>
		<div>
		<h1>SparceLinearAlgebra - Vortex</h1>
		<p>description</p>
		<hr>
		<p>our experiments:</p>
		<ul>
		{{range .}}
		<li><b>{{ .S1}}</b></li>
		<li><b>{{ .S2}}</b></li>
		{{end}}
		</ul>
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


/*
		<table>
		<thead>
		<tr>
		<th>library</th><th>date</th><th>Risc-V GPU type</th><th>description</th><th>GPU commit</th><th>Library commit</th><th>Simulator</th><th>Result</th>
		</tr>
		</thead>
		<tbody>
			<tr>
				<td> {{ .Lib}} </td>
				<td> {{ .Date}} </td>
				<td> {{ .Gpu}} </td>
				<td> {{ .Desc}} </td>
				<td> {{ .Gpucomm}} </td>
				<td> {{ .Libcomm}} </td>
				<td> {{ .Sim}} </td>
				<td> {{ .Res}} </td>
			</tr>
		</tbody>
		</table>
*/