package main

import "fmt"

func main() {
	fmt.Println(32)
	fmt.Println([]int{1, 2, 3, 4})
	// containsDuplicate([]int{11, 22, 33, 44})
	// containsDuplicate([]int{1, 2, 3, 4, 4})

	// m := map[int]bool{1: true, 2: false, 3: false}

	// fmt.Println(m)
	// fmt.Println(m[1])
	// fmt.Println(m[2])
	// m[1] = false
	// m[2] = false
	// m[3] = false
	// fmt.Println(m)
	// fmt.Println(bool(m[1]))
	// m[1] = true
	// fmt.Println(bool(m[1]))
	fmt.Println(containsDuplicate([]int{11, 22, 33, 44}))
	fmt.Println(containsDuplicate([]int{11, 22, 33, 33}))
}

func containsDuplicate(nums []int) bool {
	duplicateNumMap := make(map[int]bool)
	for _, v := range nums {
		if duplicateNumMap[v] {
			return true
		}
		duplicateNumMap[v] = true
	}
	return false
}
