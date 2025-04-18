package main

import "fmt"

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m { // 使わない場合は「_」で書いておけばOK
		fmt.Println(v)
	}

}
