package main

import "fmt"

func outer() {
	fmt.Println(111)
}
func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "hello go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "golung"
	)

	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "re go"
	fmt.Println(i3, s3)

	// 暗黙的な定義（明示的な定義と比べて、型指定の必要がない）

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 基本的には明示的な型指定をする明示的な定義を使った方が良いとされている
	// 型指定をすることでバグを抑えるように元々設計された言語なので。

	outer()

	//GOは定義された変数を必ず使う必要がある。
}
