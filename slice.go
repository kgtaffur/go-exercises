package main

import "fmt"

func main() {
	// Array
	odd := [6]int{2, 4, 6, 8, 10, 12}
	// Slice
	var s []int = odd[1:4]
	fmt.Println(s)

	// Array
	names := [4]string{
		"John",
		"Jim",
		"Jack",
		"Jen",
	}

	fmt.Println(names)

	// Slices
	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)

	slice2[0] = "ZZZ"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	// Slice literal (array literal without length)
	abc := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{6, true},
	}
	fmt.Println(abc)

	// Omit lower or upper bonds
	slice3 := []int{2, 4, 8, 10, 12, 14}
	slice4 := slice3[2:4]
	fmt.Println(slice4)

	slice5 := slice3[:3]
	fmt.Println(slice5)

	slice6 := slice3[2:]
	fmt.Println(slice6)
	fmt.Println(slice3)

	// Slice length and capacity
	sl1 := []int{2, 4, 6, 8, 10, 12, 14}
	printSlice(sl1)

	sl2 := sl1[:0] // this is like sl1[0:0]
	printSlice(sl2)

	sl3 := sl1[:4]
	printSlice(sl3)

	sl4 := sl1[2:]
	printSlice(sl4)

	// Slice make function
	slc := make([]int, 10)
	printSlice2("Slice", slc)

	slc1 := make([]int, 0, 10)
	printSlice2("Slice 1:", slc1)

	slc2 := slc1[:5]
	printSlice2("Slice 2:", slc2)

	slc3 := slc2[2:5]
	printSlice2("Slice 3:", slc3)
}

func printSlice(s []int) {
	fmt.Printf("Length = %d. Capacity = %d. | %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s -> length = %d, capacity = %d | %v\n", s, len(x), cap(x), x)
}
