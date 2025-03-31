package main

import "fmt"

// 頭文字を大文字にすると他パッケージから呼び出せる
const Pi = 3.14

const (
	URL      = "http://xxx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)
	fmt.Println(URL)
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)

}
