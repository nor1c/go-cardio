package employee

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	employee, err := GetEmployeeDetail(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		fmt.Fprintf(w, "Decode error!")
		return
	}

	// insert the data
	status, err := AddNewEmployee(&employee)
	if err != nil {
		fmt.Fprintf(w, "Error inserting record!")
		return
	}

	jsonData, err := json.Marshal(map[string]any{
		"success": status,
	})

	w.Write(jsonData)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	status, err := UpdateEmployeeDetail(&id, &employee)
	if err != nil {
		fmt.Fprintf(w, "Error updating record!")
		return
	}

	jsonData, err := json.Marshal(map[string]any{
		"success": status,
	})

	w.Write(jsonData)
}

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := DeleteEmployeeRecord(&id)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete record!")
	}

	jsonData, err := json.Marshal(map[string]any{
		"success": true,
	})

	w.Write(jsonData)
}
