package main

import "fmt"

func main() {
	x := map[string]int{"Kate": 28, "John": 37, "Raj": 20}
	fmt.Print(x)
	fmt.Println("\n", x["Raj"])

	// Insert and update
	m := make(map[string]int)
	fmt.Println(m)
	m["key1"] = 10
	m["key2"] = 20
	m["key3"] = 30
	fmt.Println(m)

	// Update
	m["key2"] = 555
	fmt.Println(m)

	// Delete
	delete(m, "key2")
	fmt.Println(m)

	// Retrieve element
	fmt.Println("The value:", m["key1"])

	// Check if present
	v, ok := m["key3"]
	fmt.Println("The value:", v, "Present?", ok)
	i, j := m["key2"]
	fmt.Println("The value:", i, "Present?", j)

	// Map of struct
	strmap := map[string]Vertex{
		"JavaTpoint": Vertex{40.68433, -7439967},
		"SSS-IT":     Vertex{37.42202, -122.08408},
	}

	fmt.Println(strmap)
}

type Vertex struct {
	Latitude, Longitude float64
}
