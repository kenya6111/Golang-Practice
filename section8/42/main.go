package main

import (
	"fmt"
)

// defer＝ 最後に実行する処理を予約）
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}
func RunDefer() {
	defer fmt.Println(11)
	defer fmt.Println(22)
	defer fmt.Println(33)
}

func main() {
	defer func() {
		if x := recover(); x != nil { // recover()はpanic状態であればxに値が入ってくる
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("start")
}
