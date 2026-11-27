package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var s []experiment

type experiment struct {
	ID string
	Lib string
	Date string
	Gpu string
	Desc string
	Gpucomm string
	Libcomm string
	Sim string
	Res string
	RepLink string
}

type test struct {
	S1 string
	S2 string
}

func h2(w http.ResponseWriter, r *http.Request) {
	s := ""
	d := template.Must(template.New("main").Parse(`<div>
		<hr>
		<h1>INFO</h1>
		<hr>
	</div>`))
	d.Execute(w, s)
}

func h(w http.ResponseWriter, r *http.Request) {
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
		<table id="Tab">
		<thead>
		<tr>
			<th>ID</th>
			<th>library</th>
			<th>date</th><th>Risc-V GPU type</th>
			<th>description</th><th>GPU commit</th>
			<th>Library commit</th>
			<th>Simulator</th>
			<th>Result</th>
			<th>Link to repository</th>
			<th>Info</th>
		</tr>
		</thead>
		<tbody>
			{{range .}}
			<tr>
				<td> {{ .ID}} </td>
				<td> {{ .Lib}} </td>
				<td> {{ .Date}} </td>
				<td> {{ .Gpu}} </td>
				<td> {{ .Desc}} </td>
				<td> {{ .Gpucomm}} </td>
				<td> {{ .Libcomm}} </td>
				<td> {{ .Sim}} </td>
				<td> {{ .Res}} </td>
				<td> {{ .RepLink}} </td>
				<td><a href="/exp{{ .ID}}"> info </a></td>
			</tr>
			{{end}}
		</tbody>
		</table>
		<a href="/info">info</a>
</div>`))
	d.Execute(w, s)
}

func main() {
	serv := http.NewServeMux()
	serv.HandleFunc("/", h)
	serv.HandleFunc("/info", h2)
	s = append(s, experiment{ID: "0", Lib: "iostream", Date: "27.11.2026", Gpu: "Nvidia GeForce RTX 5090", Desc: "test", Gpucomm: "111", Libcomm: "222", Sim: "testsim", Res: "Good", RepLink: "https://github.com"})
	s = append(s, experiment{ID: "1", Lib: "iostream1", Date: "27.11.2026", Gpu: "Nvidia GeForce RTX 5090", Desc: "test", Gpucomm: "111", Libcomm: "222", Sim: "testsim", Res: "Good", RepLink: "https://github.com"})
	s = append(s, experiment{ID: "2", Lib: "iostream2", Date: "27.11.2026", Gpu: "Nvidia GeForce RTX 5090", Desc: "test", Gpucomm: "111", Libcomm: "222", Sim: "testsim", Res: "Good", RepLink: "https://github.com"})
	serv.HandleFunc("/exp0", h2)
	serv.HandleFunc("/exp1", h2)
	serv.HandleFunc("/exp2", h2)
	err := http.ListenAndServe(":4000", serv)
	if (err != nil) {
		fmt.Println(err)
	}
}
