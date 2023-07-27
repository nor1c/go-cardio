package main

import (
	"net/http"

	employee "go-crud/modules/employee"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", employee.GetAll).Methods("GET")
	r.HandleFunc("/employees/{id}", employee.View).Methods("GET")
	r.HandleFunc("/employees", employee.AddEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", employee.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", employee.RemoveEmployee).Methods("DELETE")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err.Error())
	}
}
