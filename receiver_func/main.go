package main

import "fmt"

// struct
type User struct {
	Name string
	Age  int
}

type Admin struct {
	id   string
	pass int
}

// receiver function
func (usr User) printDetails() {
	fmt.Println("Name : ", usr.Name)
	fmt.Println("Age : ", usr.Age)
}

func (admin Admin) call(a int) {
	fmt.Println("Id : ", admin.id)
	fmt.Println("Pass : ", admin.pass)
	fmt.Println(a)
}

func main() {
	var user1 User
	var admin1 Admin

	user1 = User{
		Name: "Hasib",
		Age:  24,
	}

	admin1 = Admin{
		id:   "rcv_func",
		pass: 1234,
	}

	user1.printDetails()
	admin1.call(88)
}
