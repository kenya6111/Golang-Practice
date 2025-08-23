package main

import (
	"fmt"
	iti_1 "jituyou/1-1"
	one_2 "jituyou/1-2"
	iti_3 "jituyou/1-3"
)

var (
	a = 1
	b = 2
	c = "aaaaa"
)

const (
	aa = 11
	bb = 11
	cc = 11
)

func main() {
	iti_1.Test()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println("---------")

	one_2.TestOne_2()

	fmt.Println("---------")
	iti_3.TestIti_3()
}

func init() {
	fmt.Println("init")
}
