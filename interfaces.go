package main

import "fmt"

type vehicle interface {
	accelerate()
}

func foo(v vehicle) {
	fmt.Println(v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelerating?")
}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelerate fast?")
}

func main() {
	car1 := car{"suzuki", "blue"}
	toyota1 := toyota{"toyota", "red", 100}
	car1.accelerate()
	toyota1.accelerate()
	foo(car1)
	foo(toyota1)
}
