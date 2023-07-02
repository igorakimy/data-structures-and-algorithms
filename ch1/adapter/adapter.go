package main

import "fmt"

// IProcessor interface
type IProcessor interface {
	process()
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// process method of Adapter class
func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

// Adaptee struct
type Adaptee struct {
	adapterType int
}

// convert method of Adaptee class
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// main method
func main() {
	var processor IProcessor = Adapter{}
	processor.process()
}

// Adapter process
// Adaptee convert method
