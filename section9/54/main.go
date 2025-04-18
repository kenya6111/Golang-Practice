package main

import "fmt"

// channel
// 複数の後ルーチン間でのデータの受け渡しをするために設計されたデータ構造
func main() {
	var ch1 chan int // intのデータを保持できるチャネルを定義

	// var ch2 <-chan int // 受信専用のチャネルとして宣言
	// var ch3 chan<- int // 送信専用として定義
	ch1 = make(chan int) //チャネルの作成と初期化をして書き込みと読み込みが可能な状態にする

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	fmt.Println("----")
	ch3 <- 1              // チェネル３に一を送信するってことになる
	fmt.Println(len(ch3)) // data 1つ送ったので1になる
	ch3 <- 3
	fmt.Println(len(ch3))

	ch3 <- 4
	fmt.Println(ch3)
	fmt.Println("len:", len(ch3))

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len:", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len:", len(ch3))
	i3 := <-ch3
	fmt.Println(i3)
	fmt.Println("len:", len(ch3))

}
