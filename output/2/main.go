package main

import (
	"errors"
	"fmt"
)

func findKeyByValue(m map[int]string, val string) (int, error) {
	for k, v := range m {

		if v == val {
			return k, nil
		}
	}
	return 0, errors.New("対応するキーはありません。")
}

func main() {
	m := map[int]string{
		1: "01",
		2: "02",
		3: "03",
	}

	key, err := findKeyByValue(m, "03")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("key→%v, err→%v\n",key,err)
	}
	key2, err2 := findKeyByValue(m, "05")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("key→%v, err→%v",key2,err)
	}
}
