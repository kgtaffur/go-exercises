package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)

	argumentWithPath := os.Args
	argumentSlice := os.Args[1:]
	specificArgument := os.Args[2]

	fmt.Println(argumentWithPath)
	fmt.Println(argumentSlice)
	fmt.Println(specificArgument)
}
