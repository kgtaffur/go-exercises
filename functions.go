package main

import "fmt"

// Employee struct
type Employee struct {
	first_name string
	last_name  string
}

// Employee method
func (this Employee) fullname() {
	fmt.Println(this.first_name + " " + this.last_name)
}

// Function with return
func fun() int {
	return 12345
}

// Function with multiple return
func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0

	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}

func main() {
	// New employee
	employee1 := Employee{"John", "Ponting"}
	employee1.fullname()

	// Returned value in variable
	x := fun()
	fmt.Println(x)

	// Another function call to print values
	fmt.Println(addAll(10, 15, 20, 25, 30))
}
