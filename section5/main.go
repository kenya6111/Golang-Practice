package main

import "fmt"

// 頭文字を大文字にすると他パッケージから呼び出せる
const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const ( // 値の省略。B,Cは値を設定してないが、Aの値が入ってくれる省略の動きをする
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota // 連続する整数の連番を生成する
	c1
	c2
	c3
)

const aaa = 10

func main() {
	fmt.Println(Pi)
	fmt.Println(URL)
	fmt.Println(SiteName)
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
	fmt.Println(aaa)

}
