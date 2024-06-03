package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("LCO web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type is %T\n", res)

	Data := res.Header

	fmt.Println("Date of response is ", Data.Get("Date"))

	defer res.Body.Close() // caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("body of response", string(databyte))
}
