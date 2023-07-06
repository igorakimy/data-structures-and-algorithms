package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var template_html = template.Must(template.ParseGlob("ch2/templates/*"))

func Home(writer http.ResponseWriter, request *http.Request) {
	var customers = GetCustomers()
	log.Println(customers)
	_ = template_html.ExecuteTemplate(writer, "Home", customers)
}

func Create(writer http.ResponseWriter, request *http.Request) {
	_ = template_html.ExecuteTemplate(writer, "Create", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	customer.CustomerName = request.FormValue("customer_name")
	customer.SSN = request.FormValue("ssn")
	InsertCustomer(customer)
	var customers = GetCustomers()
	_ = template_html.ExecuteTemplate(writer, "Home", customers)
}

func Alter(writer http.ResponseWriter, request *http.Request) {
	var customer Customer
	var customerId int
	var customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = request.FormValue("customer_name")
	customer.SSN = request.FormValue("ssn")
	UpdateCustomer(customer)
	var customers = GetCustomers()
	_ = template_html.ExecuteTemplate(writer, "Home", customers)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer = GetCustomerById(customerId)
	_ = template_html.ExecuteTemplate(writer, "Update", customer)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer = GetCustomerById(customerId)
	DeleteCustomer(customer)
	var customers = GetCustomers()
	_ = template_html.ExecuteTemplate(writer, "Home", customers)
}

func View(writer http.ResponseWriter, request *http.Request) {
	var customerId int
	var customerIdStr = request.FormValue("id")
	_, _ = fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers = []Customer{customer}
	_ = template_html.ExecuteTemplate(writer, "View", customers)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	_ = http.ListenAndServe(":8000", nil)
}
