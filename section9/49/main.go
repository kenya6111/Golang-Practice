package main

import "fmt"

func main() {
	sl := []string{"a", "b", "c"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	for i := range sl {
		fmt.Println(i)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}
