package main

import (
	"fmt"
	"strconv"
)

func main() {
	sl := []interface{}{1, "2", 10, "11", "abc", true}
	for _, v := range sl {
		switch value := v.(type) {
		case int:
			fmt.Printf("%02d\n", v)
		case string:
			if i, err := strconv.Atoi(value); err == nil {
				fmt.Printf("%02d\n", i)
			} else {

				fmt.Printf("変換エラー: '%s' は数値ではありません\n", v)
			}
		default:
			fmt.Printf("データ型エラー: %Tは未対応のデータ型です\n", v)

		}
	}
}
