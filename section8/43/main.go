package main

import (
	"fmt"
	"time"
)

// gorutin

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	// go文を使うことで並行処理を簡単に実行できる
	go sub()

	for {
		fmt.Println("main Loop")
		time.Sleep(200 * time.Millisecond)
	}

}
