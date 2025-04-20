package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 20},
		3: {Name: "user2", Age: 30},
	}

	fmt.Println(m)
	m2 := map[User]int{
		{Name: "user1", Age: 20}: 13,
		{Name: "user2", Age: 30}: 22,
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	m3[11] = User{Name: "user22", Age: 22}
	m3[113] = User{Name: "user23", Age: 23}

	fmt.Println(m3)

	for _, v := range m3 {
		fmt.Println(v)
	}
}
