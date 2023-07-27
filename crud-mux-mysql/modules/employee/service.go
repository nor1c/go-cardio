package employee

import (
	"database/sql"
	"go-crud/infra/db"
)

func connect() *sql.DB {
	conn := db.Connect()

	return conn
}

func query(query string) (*sql.Rows, error) {
	conn := connect()

	result, err := conn.Query(query)
	defer conn.Close()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllEmployees() ([]Employee, error) {
	q, err := query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer q.Close()

	employees := []Employee{}

	for q.Next() {
		var id int
		var name, city string

		err = q.Scan(&id, &name, &city)
		if err != nil {
			return nil, err
		}

		emp := Employee{
			Id:   id,
			Name: name,
			City: city,
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

func GetEmployeeDetail(id string) (*Employee, error) {
	qRes, err := query("SELECT * FROM employee WHERE id=" + id)
	if err != nil {
		return nil, err
	}
	defer qRes.Close()

	employee := &Employee{}

	for qRes.Next() {
		err := qRes.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			return nil, err
		}
	}

	return employee, nil
}

func AddNewEmployee(employee *Employee) (bool, error) {
	conn := connect()
	defer conn.Close()

	newEmp, err := conn.Prepare("INSERT INTO employee(name, city) VALUES (?,?)")
	if err != nil {
		return false, err
	}
	newEmp.Exec(employee.Name, employee.City)
	defer newEmp.Close()

	return true, nil
}

func UpdateEmployeeDetail(id *string, employee *Employee) (bool, error) {
	conn := connect()
	defer conn.Close()

	update, err := conn.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
	if err != nil {
		return false, err
	}
	update.Exec(employee.Name, employee.City, id)
	defer update.Close()

	return true, nil
}

func DeleteEmployeeRecord(id *string) (bool, error) {
	_, err := query("DELETE FROM employee WHERE id=" + *id)
	if err != nil {
		return false, err
	}

	return true, nil
}
