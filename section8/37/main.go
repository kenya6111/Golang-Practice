package main

import (
	"fmt"
)

func main() {
	fmt.Println(11)

	arr := [3]string{"aaa", "bbb", "ccc"}

	// 範囲式forと呼ぶ
	for i, v := range arr {
		fmt.Println(i, v)
	}

	m := map[int]string{1: "aaa", 2: "bbb", 3: "ccc"}

	fmt.Println(m)
}
