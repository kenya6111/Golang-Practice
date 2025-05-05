package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	for i, _ := range users {
		fmt.Println(i)
		users[i].Age = 44
	}

	fmt.Printf("%v", users) // どうなる？

	// user は users の各要素のコピーにすぎず、user.Age = 44 はコピー側の Age を変更してるだけなため
}
