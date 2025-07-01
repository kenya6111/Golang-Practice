package main

import "fmt"


func main(){
	var ch1 chan int // 特にサブタイプを指定しないと双方間の通信が可能

	// var ch2 <-chan int // 受信専用のチャネル

	// var ch3 chan<- int//送信専用



	ch1 = make(chan int) //チャネルの生成と初期化ができ、書き込み可能な状態になる

	ch2 := make(chan int)// このように型推論を使って定義するのが普通。

	fmt.Println(cap(ch1))// チャネルの容量＝バッファ 0
	fmt.Println(cap(ch2))// チャネルの容量＝バッファ 0

	ch3 := make(chan int,5)
	fmt.Println(cap(ch3))// 5

	//データの送信
	ch3 <-1 // ch3に1のデータを送信
	fmt.Println(len(ch3))// データを1つ送ったので1が出力される

	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3))// データを合計3つ送ったので2が出力される
	
	i:= <-ch3// チャネルからデータを取り出す
	fmt.Println(i)// 1が出力される
	fmt.Println(len(ch3))// データを1個取り出したので2が出力される
	
	i2 := <-ch3
	fmt.Println(i2)// 2が出力される
	fmt.Println(<-ch3)// 3が出力される
	fmt.Println(len(ch3))// データを2個取り出したので0が出力される
	
	// バッファサイズを超えたデータを送るとどうなるのか
	
	ch3 <- 10
	ch3 <- 20
	ch3 <- 30
	ch3 <- 40
	ch3 <- 50
	fmt.Println(len(ch3))
	ch3 <- 60 // バッファサイズを超えたデータをおこうるとデットロックが起こる。
	// このように、複数のデータを保持する場合はバッファサイズを指定する必要がある


}


// チャネル→複数のGoルーチン間でのデータの受け渡しをするために設計されたデータ構造

// サブタイプを指定して受信専用、送信専用とかのチャネルを生成できる