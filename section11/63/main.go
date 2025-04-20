package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 構造体のメソッドは、任意の型に特化した関数を定義するための仕組み
func (u User) sayName() { // メソッド名の前にレシーバを書く
	fmt.Println(u.Name)
}
func (uu User) outputName() { // メソッド名の前にレシーバを書く
	fmt.Println(uu.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}
func (u *User) SetName2(name string) { // レシーバをポインタ型にしているので参照私できるようになった
	u.Name = name
}
func main() {
	user1 := User{Name: "user1"}
	user1.sayName()

	user1.SetName("A")
	user1.sayName()

	user1.SetName2("A")
	user1.sayName()

	user2 := User{Name: "user2"}
	user2.sayName()
	user2.outputName()
}
