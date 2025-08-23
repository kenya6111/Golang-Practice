package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(111)
	v := true
	s := strconv.FormatBool(v)
	fmt.Printf("%T, %v\n", s, s)

	v2 := int64(-42)

	s10 := strconv.FormatInt(v2, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v2, 16)
	fmt.Printf("%T, %v\n", s16, s16)

	v3 := strconv.Itoa(44)
	fmt.Printf("%T, %v\n", v3, v3)

}
