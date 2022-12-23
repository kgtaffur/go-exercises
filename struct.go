package main

import "fmt"

// This is a struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Go embedded struct
type Employee struct {
	Person
	employeeId int
}

// Person struct "method"
func (this Person) details() {
	fmt.Println(this, "I am a person")
}

// Employee struct "method"
func (this Employee) details() {
	fmt.Println(this, "I am a employee")
}

func main() {
	x := Person{age: 30, firstName: "John", lastName: "Anderson"}
	fmt.Println(x)
	fmt.Println(x.firstName)

	person1 := Person{"Raj", "Kumar", 25}
	person1.details()

	employee1 := Employee{Person{"John", "Ponting", 24}, 11}
	employee1.details()
}
