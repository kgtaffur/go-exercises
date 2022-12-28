package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 0 to 100
	fmt.Println(rand.Intn(100))

	// 0 to 1
	fmt.Println(rand.Float64())

	rand.Seed(time.Now().Unix())
	myrand := random(1, 20)

	fmt.Println(myrand)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
