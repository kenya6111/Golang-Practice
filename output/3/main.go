package main

import (
	"fmt"
	"slices"
)

type MyIntSlice []int

func (Mis MyIntSlice) Unique() []int {
	res := []int{}
	for _, v := range Mis {
		if !slices.Contains(res, v) {
			res = append(res, v)
		}
	}
	Mis = res
	return Mis
}
func main() {
	m := MyIntSlice{1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(m.Unique()) // [1, 2, 3, 4, 5]
}
