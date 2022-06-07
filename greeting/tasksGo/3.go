package main

import "fmt"

type Employee struct {
	Name   string
	Rate   int
	Active bool
	Address
}
type Subscriber struct {
	Name        string
	Age         int
	HomeAddress Address
}
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func main() {
	address := Address{"Costin", "Chisinau", "Moldova", "0271"}
	employee1 := Employee{Name: "Valera", Rate: 25, Active: false}
	employee1.Address = address

	employee1.City = "Moscow"

	subscriber := Subscriber{Name: "vova", Age: 25}
	subscriber.HomeAddress.Street = "Chisinau"
	fmt.Println(employee1, subscriber)
}
