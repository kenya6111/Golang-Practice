package main

import (
	"fmt"
	"unicode/utf8"
)

type T struct {
	User User // フィールド名と型が同じなのはよくある
}
type User struct { // 構造体はフィールドに構造体を埋め込める
	Name string
	Age  int
}

// 基本型に名前をつける（≒別の型にする）
type MyInt int

// 構造体（struct）を定義する
type User2 struct {
	Name string
	Age  int
}

// Rune は int32 のエイリアス（完全に同じ型）
type Rune = int32

func main() {
	t := T{
		User: User{Name: "user1", Age: 10},
	}
	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	var a MyInt = 1
	fmt.Println(a)
	// var b int = a エラー！型が違う

	u := User2{Name: "kenya", Age: 25}
	fmt.Println(u)

	var r Rune = '世'
	fmt.Println(r)        // => 19990（Unicodeコードポイント）
	fmt.Printf("%c\n", r) // => 世（文字として出力）

	// utf8パッケージなどの関数も int32 としてそのまま使える
	size := utf8.RuneLen(r)
	fmt.Println("バイト数:", size) // => バイト数: 3

}
