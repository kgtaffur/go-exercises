package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// String
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	// Length
	str := "I love my country"
	fmt.Println(len(str))

	// ASCII code
	fmt.Println("ASCII value of A is:", "A"[0])

	// Upper
	strg := "india"
	fmt.Println(strings.ToUpper(strg))

	// Lower
	strg2 := "INDIA"
	fmt.Println(strings.ToLower(strg2))

	// Has prefix
	s := "INDIA"
	fmt.Println(strings.HasPrefix(s, "IN"))

	// Has suffix
	fmt.Println(strings.HasSuffix(s, "IA"))

	// Join
	var arr = []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(arr, "*"))

	// Repeat
	var st = "New "
	fmt.Println(strings.Repeat(st, 4))

	// Contains
	text := "Hi... there"
	fmt.Println(strings.Contains(text, "th"))

	// Index
	fmt.Println(strings.Index(text, "th"))

	// Count
	fmt.Println(strings.Count(text, "e"))

	// Replace
	fmt.Println(strings.Replace(text, "e", "Z", 2))

	// Split
	sentence := "I,love,my,country"
	var arr2 []string = strings.Split(sentence, ",")
	fmt.Println(len(arr2))
	for i, v := range arr2 {
		fmt.Println("Index:", i, "value:", v)
	}

	fmt.Printf("%q\n", strings.Split("x,y,x", ","))
	fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn", "and"))
	fmt.Printf("%q\n", strings.Split(" abc", ""))
	fmt.Printf("%q\n", strings.Split("", "Hello"))

	// Compare
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	// Trim
	fmt.Println(strings.TrimSpace(" \t\n I love my country \n\t\r\n"))

	// Contains any
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))
}
