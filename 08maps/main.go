package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps section ")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["CPP"] = "c++"
	languages["PY"] = "Python"
	languages["RB"] = "RUBY"

	fmt.Println("List of all languages: ", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages,"RB")


	// loops 
	for key,value := range languages {
		fmt.Printf("For key %v, value is %v \n",key,value )
	}
} 