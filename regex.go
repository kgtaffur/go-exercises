package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Find string
	re := regexp.MustCompile(".com")
	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("abc.org"))
	fmt.Println(re.FindString("fb.com"))

	// Find string index
	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("abc.org"))
	fmt.Println(re.FindStringIndex("fb.com"))

	// Find string submatch
	re2 := regexp.MustCompile("f([a-z]+)ing")
	fmt.Println(re2.FindStringSubmatch("flying"))
	fmt.Println(re2.FindStringSubmatch("abcfloatingxyz"))
}
