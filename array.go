package main

import "fmt"

func main() {
	// Array
	var x [5]int
	var i, j int
	for i = 0; i < 5; i++ {
		x[i] = i + 10
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, x[j])
	}

	// Multidimensional array
	var a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var i2, j2 int
	for i2 = 0; i2 < 3; i2++ {
		for j2 = 0; j2 < 3; j2++ {
			fmt.Print(a[i2][j2])
		}
		fmt.Println()
	}
}
