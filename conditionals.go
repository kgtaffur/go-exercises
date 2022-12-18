package main

import "fmt"

func main() {
	// Go if
	var a int = 10
	if a%2 == 0 {
		fmt.Printf("a is even number\n")
	}

	// Go if-else
	var b int = 10
	if b%2 == 0 {
		fmt.Printf("b is even\n")
	} else {
		fmt.Printf("b is odd\n")
	}
	fmt.Printf("value of b is: %d\n", b)

	// Go if-else with input from user
	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)

	if input%2 == 0 {
		fmt.Printf(" is even\n")
	} else {
		fmt.Printf(" is odd\n")
	}

	// Go if else-if
	fmt.Print("Enter text: ")
	var input2 int
	fmt.Scanln(&input2)
	if input2 < 0 || input2 > 100 {
		fmt.Print("Please enter a valid number")
	} else if input2 >= 0 && input2 < 50 {
		fmt.Print(" Fail\n")
	} else if input2 >= 50 && input2 < 60 {
		fmt.Print(" D grade\n")
	} else if input2 >= 60 && input2 < 70 {
		fmt.Print(" C grade\n")
	} else if input2 >= 70 && input2 < 80 {
		fmt.Print(" B grade\n")
	} else if input2 >= 80 && input2 < 90 {
		fmt.Print(" A grade\n")
	} else if input2 >= 90 && input2 <= 100 {
		fmt.Print(" A+ grade\n")
	}

	// Go nested if-else
	var x int = 10
	var y int = 20
	if x >= 10 {
		if y >= 10 {
			fmt.Printf("Inside nested if statement\n")
		}
	}
	fmt.Printf("Value of x is: %d\n", x)
	fmt.Printf("Value of y is: %d\n", y)
}
