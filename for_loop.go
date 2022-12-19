package main

import "fmt"

func main() {
	// Go for loop, counter-controll
	for a := 0; a < 11; a++ {
		fmt.Print(a, "\n")
	}

	// Go for loop, conditional-controlled iteration
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}
