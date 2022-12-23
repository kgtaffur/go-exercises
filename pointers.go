package main

import "fmt"

func main() {
	x := 10
	changeX(&x)
	fmt.Println(x)

	// new() returns a pointer of the specified type
	ptr := new(int)
	fmt.Println("Before change ptr", *ptr)
	changeX(ptr)
	fmt.Println("After change ptr", *ptr)
}

func changeX(x *int) {
	*x = 10
}
