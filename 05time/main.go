package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	presentTime := time.Now();
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createDate := time.Date(2024,time.April,4,23,25,0,0,time.UTC);
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}