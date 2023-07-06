package main

import "database/sql"

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

func GetCustomerById(customerID int) Customer {
	var database = GetConnection()

	rows, err := database.Query("SELECT * FROM Customer WHERE CustomerId=?", customerID)
	if err != nil {
		panic(err.Error())
	}
	var customer = Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var SSN string
		if err = rows.Scan(&customerId, &customerName, &SSN); err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = SSN
	}
	defer database.Close()
	return customer
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

}
