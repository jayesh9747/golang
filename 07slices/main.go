package main

import (
	"fmt"
	"sort"
)

func main() {

	// initialization of slices
	
	var list = []string{}

	fmt.Println(list)

	fmt.Println("Welcome to the slice in golang")

	var fruitList = []string{"Apple","Orange"}

	fmt.Printf("Type of fruitList is %T\n", fruitList)


	//adding or pushing data in the slice 
	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

 
	vegList := fruitList[1:]

	fmt.Println(vegList)

	highScores := make([]int , 4)

	highScores[0] = 100
	highScores[1] = 900
	highScores[2] = 700
	highScores[3] = 800

	//highScore[4] = 900  we don't do this cause of the it allocate only four integer space 

	// it reallocating the memory to the slice
	highScores = append(highScores, 500,600,700)

	fmt.Println(highScores)

	//sort function
	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	// how to remove a value from slices based on index 

	courses := []string{"reactjs","javascript","swift","python","ruby"}

	fmt.Println(courses)

	index :=2

	// use append to remove elements from slice 
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
	


}