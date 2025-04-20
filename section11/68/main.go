package main

import "fmt"

type MyInt int

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)
}
