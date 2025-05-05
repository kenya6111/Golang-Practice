package main

import "fmt"

type MyInt int

func (m MyInt) String() string {
	return "hoge"
}

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}
func main() {
	var m MyInt = 3
	fmt.Println(m) // hogeと出力させるように修正せよ。ただしmain関数に変更を加えないこと。
}
