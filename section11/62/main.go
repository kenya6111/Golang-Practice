package main

import "fmt"

type User struct { // これで構造体の定義ができる Goの struct は「TypeScriptの型定義」に近い
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "TEST"
	user.Age = 9999
}
func main() {
	// 構造体はクラスのような存在。複数の任意の型の値を1つにまとめたもの。

	var user1 User

	fmt.Println(user1)

	user1.Name = "sada"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 100
	fmt.Println(user2)

	user3 := User{Name: "user3"}
	fmt.Println(user3)
	user4 := User{"user4", 40}
	fmt.Println(user4)
	user6 := User{Name: "usr6"}
	fmt.Println(user6)

	fmt.Println("---")
	user7 := new(User) // newで定義した変数は構造体のポインタ型をかえす
	fmt.Println(user7)
	fmt.Println(*user7)

	fmt.Println("---")
	user8 := &User{Name: "aaa"} // アドレス演算子でもポインタ型で宣言できる
	fmt.Println(user8)
	fmt.Println(*user8)
	fmt.Println("---")

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}
