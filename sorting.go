package main

import (
	"fmt"
	"sort"
)

func main() {
	intValue := []int{10, 20, 5, 8}
	sort.Ints(intValue)
	fmt.Println("Ints:", intValue)

	floatValue := []float64{10.5, 20.5, 5.5, 8.5}
	sort.Float64s(floatValue)
	fmt.Println("Floats:", floatValue)

	stringValue := []string{"Raj", "Mohan", "Roy"}
	sort.Strings(stringValue)
	fmt.Println("Strings:", stringValue)

	str := sort.Float64sAreSorted(floatValue)
	fmt.Println("Float sorted:", str)

	cities := []string{"New York", "London", "Washington", "Delhi"}
	sort.Sort(OrderByLengthDesc(cities))
	fmt.Println(cities)
}

type OrderByLengthDesc []string

func (this OrderByLengthDesc) Len() int {
	return len(this)
}
func (this OrderByLengthDesc) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this OrderByLengthDesc) Less(i, j int) bool {
	return len(this[i]) > len(this[j])
}
