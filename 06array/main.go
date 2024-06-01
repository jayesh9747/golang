package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "orange"
	fruitList[3] = "pineapple"
	 

	fmt.Println(fruitList)

	for i := 0 ; i < len(fruitList) ; i++ {
		fmt.Println(fruitList[i])
	}

	var vegList = [3]string{"potato"}

	studentList := [4]string{"Jayesh","Mahesh","Ramesh"}

	studentList[3] = "raman"
	
	fmt.Println(vegList)
	fmt.Println(studentList)

}