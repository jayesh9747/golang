package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang");


	days := []string{"Sunday", "Tuesday", "Friday"}

	fmt.Println(days)

	// for i :=0 ; i < len(days); i++{
	// 	fmt.Println(days[i])
	// }

	for j := range(days){
		fmt.Println(days[j]);
	}

	for i , day := range days{
		fmt.Printf("index is %v and value is %v\n",i,day)
	}

	

	i:= 2
	if(i==2){
		goto lco
	}	

	lco:
		fmt.Println("you should jump tp the lco web")
		
		


}
