package main

import "fmt"

var i11 int = 10

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
	//:= は、「短変数宣言 (short variable declaration)」 と呼ばれる構文です。
	// 変数の型を明示的に指定せずに、新しい変数を宣言しながら初期化する際に使われます。

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 基本的には明示的な型指定をする明示的な定義を使った方が良いとされている
	// 型指定をすることでバグを抑えるように元々設計された言語なので。

	x, y := "aaa", 123
	fmt.Print(x, y)
	fmt.Print("\n")
	outer()

	fmt.Print("\n")
	var aaa string
	fmt.Print("\n")
	var bbb int
	fmt.Print(aaa)
	fmt.Print(bbb)

	fmt.Print("\n")
	i10 := 300
	fmt.Print(i10)
	fmt.Print("\n")
	fmt.Printf("i10=%T", i10)

	i10 = 111
	fmt.Print(i11)

	// 	使い分け
	// := を使う場合
	//   関数の中でローカル変数を素早く宣言・初期化したいときに便利。
	//   型推論でスッキリ書ける。

	// var を使う場合
	//   関数外(パッケージスコープ) では := が使えないので必須。
	//   型を明示したい場合、またはパッケージのレベルで複数の変数をまとめて定義したい場合など。
	// 	:= は「短変数宣言」 で、ローカル変数を新規に宣言・初期化するときに使う。

	// 少なくとも1つは新しい変数が必要 (既存の変数しかない場合はエラー)。

	// パッケージスコープ(グローバル)では使えず、関数の中でのみ使用可能。
}
