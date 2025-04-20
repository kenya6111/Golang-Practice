package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {

	user1 := NewUser("user1", 25)
	fmt.Println(user1)
	fmt.Println(&user1)

	user2 := NewUser("user2", 26)
	fmt.Println(user2)
	fmt.Println(&user2)

	user3 := NewUser("user3", 27)
	fmt.Println(user3)
	fmt.Println(&user3)
}
