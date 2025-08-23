package iti_3

import "fmt"

const (
	X = iota + 1 // 1
	Y            // 2
	Z            // 3
)

func TestIti_3() {
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
}
