package main

import (
	"log"
	"time"
)

func haveFun(s string) {
	log.Printf("\tA: Let's have fun: %v", s)
}

func doPolling() {
	for _ = range time.Tick(2 * time.Second) {
		haveFun("\tB:Okay!")
	}
}

func main() {
	go doPolling()
	select {}
}
