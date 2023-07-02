package main

import (
	"fmt"
)

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

type DataTransferObject interface {
	getId() string
}

type Customer struct {
	id   string
	name string
	ssn  string
}

func (customer Customer) getId() string {
	fmt.Println("getting customer Id")
	return customer.id
}

type Employee struct {
	id   string
	name string
}

func (employee Employee) getId() string {
	fmt.Println("getting employee Id")
	return employee.id
}

type Manager struct {
	id   string
	name string
	dept string
}

func (manager Manager) getId() string {
	fmt.Println("getting manager Id")
	return manager.id
}

type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (address Address) getId() string {
	fmt.Println("getting address Id")
	return address.id
}

func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}
	var customer = factory.getDataTransferObject("customer")
	fmt.Println(customer.getId())
	var employee = factory.getDataTransferObject("employee")
	fmt.Println(employee.getId())
	var manager = factory.getDataTransferObject("manager")
	fmt.Println(manager.getId())
	var address = factory.getDataTransferObject("address")
	fmt.Println(address.getId())
}

// new DTO of dtoType: customer
// getting customer Id
// 1
// new DTO of dtoType: employee
// getting employee Id
// 2
// new DTO of dtoType: manager
// getting manager Id
// 3
// new DTO of dtoType: address
// getting address Id
// 4
