package main

import "fmt"

func main() {

	fmt.Println("Welcome to the slice in golang")

	var fruitList = []string{"Apple","Orange"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)


	vegList := fruitList[1:]

	fmt.Println(vegList)



}