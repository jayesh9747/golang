package main

import "fmt"

func main() {

	fmt.Println("welcome to Struct in Golang")
	//no inheritance in golang, No super or parent class


	jayesh := User{"Jayesh","jayeshsavaliya@gmail.com",true,18}

	fmt.Println(jayesh)

	fmt.Printf("jayesh details are: %+v\n",jayesh)

}

// creating struct
// all name is capital must aware of this 
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
