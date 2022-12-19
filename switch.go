package main

import "fmt"

func main() {
	fmt.Print("Enter number: ")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 10:
		fmt.Print("the value is 10\n")
	case 20:
		fmt.Print("the value is 20\n")
	case 30:
		fmt.Print("the value is 30\n")
	case 40:
		fmt.Print("the value is 40\n")
	default:
		fmt.Print("It is not 10, 30, 40\n")
	}

	k := 30
	switch k {
	case 10:
		fmt.Println("was <= 10")
		fallthrough
	case 20:
		fmt.Println("was <= 20")
		fallthrough
	case 30:
		fmt.Println("was <= 30")
		fallthrough
	case 40:
		fmt.Println("was <= 40")
		fallthrough
	case 50:
		fmt.Println("was <= 50")
		fallthrough
	case 60:
		fmt.Println("was <= 60")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
