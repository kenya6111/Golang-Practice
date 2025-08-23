package main

import "fmt"

type MyInt int

func main() {
	my := MyInt(1)
	fmt.Println(my)
	fmt.Printf("%T\n", my)
}

func (i MyInt) test() {
	fmt.Println(1111)
}
