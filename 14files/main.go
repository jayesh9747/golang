package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to filed in Golang")
	content := "This needs to go in a file- LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	checknilErr(err)

	length, err := io.WriteString(file,content)

	checknilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checknilErr(err)

	fmt.Println("Text data inside the file is\n", string(databyte))

}

func checknilErr(e error) {
	if e != nil {
		panic(e)
	}
}
