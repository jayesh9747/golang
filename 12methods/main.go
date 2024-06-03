package main

import "fmt"

func main() {

    fmt.Println("Welcome to Struct in Golang")
    // No inheritance in Golang, No super or parent class

    jayesh := User{"Jayesh", "jayeshsavaliya@gmail.com", true, 18}

    fmt.Println(jayesh)
    fmt.Printf("Jayesh details are: %+v\n", jayesh)

    jayesh.GetStatus()

    fmt.Printf("Before: %+v\n", jayesh)
    
    jayesh.SetName()

    fmt.Printf("After: %+v\n", jayesh)
}

// Creating struct
// Note: All field names must start with an uppercase letter to be exported
type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}

func (u User) GetStatus() {
    fmt.Println("Is user active:", u.Status)
}

// Here we change the Name of user so we pass object by reference 
func (u *User) SetName() {
    u.Name = "Jayesh Savaliya"
    fmt.Println("User Name is", u.Name)
}
