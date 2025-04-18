package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println(1 + 2)
	fmt.Println(1 - 2)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)
	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)

	fmt.Println(1 == 1)
	fmt.Println(1 == 3)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false || false)
	fmt.Println(false || true)
	fmt.Println(true || true)
	fmt.Println(!true)
	fmt.Println(!false) // !で論理値を反転させることができる

}
