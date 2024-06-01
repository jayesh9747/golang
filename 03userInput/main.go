package main

import "fmt"

func main() {
	welcome:= "Welcome to User Input"
	fmt.Println(welcome);

	var UserInput string;
	
	fmt.Scanln(&UserInput)

	fmt.Println("User Input is the", UserInput)

}