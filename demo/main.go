package main

import (
	"fmt"
	"log"
)

// Method to add two numbers
func add(a, b int) int {
	log.Println("Adding", a, b)
	return a + b
}

// Method to subtract two numbers
func sub(a, b int) int {
	log.Println("Subtracting", a, b)
	return a - b
}

func main() {
	log.Println("Adding", 1, 2)
	b := add(1, 2)
	log.Println("Addition is", b)
	fmt.Println("Done with execution")
}
