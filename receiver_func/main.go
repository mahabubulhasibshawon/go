package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// receiver function
func (usr User) printDetails() {
	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Hasib",
		Age:  24,
	}

	user1.printDetails()
}
