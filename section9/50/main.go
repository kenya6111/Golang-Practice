package main

import "fmt"

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	fmt.Println(sum(1, 2, 3, 4))

	sl := []int{1, 2, 3}
	fmt.Println(sum(sl...))

	sl3 := []int{1, 2, 3, 4, 5, 6}

	num := 0
	for _, v := range sl3 {
		num = num + v
	}
	fmt.Println(num)

	sl4 := []int{8, 8, 8, 8, 8}

	fmt.Println(sum(sl4...))

}
