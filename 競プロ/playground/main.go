package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())

	a := [3]string{
		"aa",
		"bb",
		"cc",
	}
	fmt.Println(a)
	fmt.Println(strconv.Atoi("22"))
	i, err := strconv.Atoi("-42")
	fmt.Println(i)
	fmt.Println(err)
}
