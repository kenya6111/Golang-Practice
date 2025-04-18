package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)
	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 300, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))

	fmt.Println(cap(sl2))
	fmt.Println("----")
	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

}
