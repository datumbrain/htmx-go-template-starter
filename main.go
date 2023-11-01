package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Employee struct {
	Name       string
	Age        int
	EmployeeID string
}

var emps = map[string][]Employee{
	"Managers": {
		{Name: "John Doe", Age: 40, EmployeeID: "EMP001"},
		{Name: "Jane Doe", Age: 35, EmployeeID: "EMP002"},
	},
}

func appendEmps(emps *map[string][]Employee, dept string, emp Employee) bool {
	if _, ok := (*emps)[dept]; ok {
		(*emps)[dept] = append((*emps)[dept], emp)
		return true
	}
	return false
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, emps)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		name := r.PostFormValue("name")
		age, _ := strconv.Atoi(r.PostFormValue("age"))

		appendEmps(&emps, "Managers", Employee{Name: name, Age: age, EmployeeID: "EMP00x"})

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "emp-list-element", Employee{Name: name, Age: age, EmployeeID: "EMP00x"})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-emp", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
