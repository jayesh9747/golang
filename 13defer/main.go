package main

import "fmt"

func main() {
	// Last in first out ->  Stack
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}


// Two , One , World  