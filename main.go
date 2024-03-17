package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5433"
	user     = "root"
	password = "password"
	dbname   = "pgsql"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s  sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfuly connected to database")

	//CreateEmployee()
	//GetEmplyes()
	//UpdateEmployee()
	DeleteEmployee()
}

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `INSERT INTO employees (full_name, email,  age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`
	err = db.QueryRow(sqlStatement, "ramdan", "ramdan@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employe Data: %+v\n", employee)
}

func GetEmplyes() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email=$3, division=$4, age=$5
	WHERE id=$1;`

	res, err := db.Exec(sqlStatement, 1, "Dinda", "dindakusuma@gmail.com", "MARKETING", 22)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE FROM employees
	WHERE id=$1;`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete data amount:", count)
}
