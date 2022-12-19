package main

import "fmt"

func main() {
	var input int
Loop:
	fmt.Println("You are not elegible to vote	")
	fmt.Print("Enter your age: ")
	fmt.Scanln(&input)
	if input <= 17 {
		goto Loop
	} else {
		fmt.Println("You can vote")
	}
}
