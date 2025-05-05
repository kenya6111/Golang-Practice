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

func (u User) saySay() {
	fmt.Println("say!!")
}
func (uu User) outputName() { // メソッド名の前に「「レシーバ」」を書く
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
	user1.saySay()
	fmt.Println("---")

	user1.SetName("A")
	user1.sayName()
	fmt.Println("---")

	user1.SetName2("A")
	user1.sayName()

	user2 := User{Name: "user2"}
	user2.sayName()
	user2.outputName()
}

// []int（スライス） は 内部的にポインタ＋長さ＋容量を持つ構造体のようなもの。
// なので、値渡しでも中身（配列の要素）は変更可能。

// 構造体は値型
// func (u User) SetName(name string) {
// 	u.Name = name // これはコピーに対する操作 → 元の user1 は変わらない
// }
// func (u *User) SetName2(name string) {
// 	u.Name = name // ポインタ経由 → 元の user1.Name を直接変更できる
// }
// User は 値型なので、メソッドで値レシーバを使うとコピーされる。
// 値を変えたいなら、ポインタレシーバ (*User) にする必要がある。
