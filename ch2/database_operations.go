package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() *sql.DB {
	database, err := sql.Open("sqlite3", "./ch2/crm")
	if err != nil {
		panic(err.Error())
	}
	return database
}

func GetCustomers() []Customer {
	var database = GetConnection()

	rows, err := database.Query("SELECT * FROM Customer ORDER BY CustomerId DESC")
	if err != nil {
		panic(err.Error())
	}
	var customer = Customer{}

	var customers []Customer
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		if err = rows.Scan(&customerId, &customerName, &ssn); err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()

	return customers
}

func InsertCustomer(customer Customer) {
	var database = GetConnection()

	insert, err := database.Prepare("INSERT INTO Customer(CustomerName, SSN) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, _ = insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}

func UpdateCustomer(customer Customer) {
	var database = GetConnection()

	update, err := database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if err != nil {
		panic(err.Error())
	}
	_, _ = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	defer database.Close()
}

func DeleteCustomer(customer Customer) {
	var database = GetConnection()

	del, err := database.Prepare("DELETE FROM Customer WHERE CustomerId=?")
	if err != nil {
		panic(err.Error())
	}
	_, _ = del.Exec(customer.CustomerId)
	defer database.Close()
}

func main() {
	// var customers = GetCustomers()

	// fmt.Println("Before Insert", customers)
	// var customer Customer
	// customer.CustomerName = "Arnie Smith"
	// customer.SSN = "2386343"
	// InsertCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Insert", customers)

	// fmt.Println("Before Update", customers)
	// var customer Customer
	// customer.CustomerName = "George Thompson"
	// customer.SSN = "23233432"
	// customer.CustomerId = 5
	// UpdateCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Update")

	// fmt.Println("Before Delete", customers)
	// var customer Customer
	// customer.CustomerName = "George Thompson"
	// customer.SSN = "23233432"
	// customer.CustomerId = 5
	// DeleteCustomer(customer)
	// customers = GetCustomers()
	// fmt.Println("After Delete", customers)
}
