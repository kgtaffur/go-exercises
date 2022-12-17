package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v: %T\n", i, i)
	fmt.Printf("%v: %T\n", f, f)
	fmt.Printf("%v: %T\n", b, b)
	fmt.Printf("%q: %T\n", s, s)
}
