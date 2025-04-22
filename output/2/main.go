package main

import (
	"errors"
	"fmt"
)

func findKeyByValue(m map[int]string, val string) (int, error) {
	for k, v := range m {

		if val == v {
			return k, nil
		}
	}
	return 0, errors.New("error!!!!!")
}
func main() {
	m := map[int]string{
		1: "01",
		2: "02",
		3: "03",
	}
	// fmt.Println(findKeyByValue(m, "01"))

	key, err := findKeyByValue(m, "03") // key→3, err→nil
	fmt.Println("key", key)
	fmt.Println("err", err)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key)
	}
	key2, err2 := findKeyByValue(m, "05") // key→0にすること(初期値なので), errはある
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(key2)
	}
}
