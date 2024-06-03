package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	greeter()


	result := adder(3,5)

	fmt.Println("Result is: ", result)

	ansPro := proAdder(1,2,3,4,5,6)

	fmt.Println("Result of proAns",ansPro)

	todo := []string{"t1","t2","t3"}

	fmt.Printf("%v\n", todo)
}

func adder(a int , b int) int {
	return a + b 
}

func greeter(){
	fmt.Println("Welcome my friends!")
}


func proAdder(values ...int) int{
	ans := 0
	for i := range(values){
		ans+= values[i]
	}
	return ans
}



// func T([]string t){

// }