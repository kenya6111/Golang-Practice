package main

import "fmt"

type T struct {
	User User // フィールド名と型が同じなのはよくある
}
type User struct { // 構造体はフィールドに構造体を埋め込める
	Name string
	Age  int
}

func main() {
	t := T{
		User: User{Name: "user1", Age: 10},
	}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
}
