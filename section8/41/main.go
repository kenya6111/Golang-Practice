package main

import (
	"fmt"
	"os"
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
	TestDefer()
	// defer func() {
	// 	fmt.Println(1)
	// 	fmt.Println(2)
	// 	fmt.Println(3)
	// }()

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("hello"))
}
