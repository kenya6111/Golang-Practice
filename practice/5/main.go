package main

import (
	"fmt"
	"time"
)


func reciever (ch chan int){
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main(){
	// ch1 := make(chan int)
	// fmt.Println(<-ch1)
	// ここでデッドロックとなる。Goのチャネルは単なるキューではなく、Goルーチン間のデータの共有のためにある。
	// チャネルから受信するは、裏を返せば他のごルーチンがチャネルをデータ送信するのを待つってこと
	// ここではmain関数を処理するための語ルーチンしかそもそもない＝ch1にデータを送信してくれるGoルーチンがそもそも存在しないのでデッドロックになる
	// 何回も優雅ちゃねるはGoルーチン間でデータを共有するために仕組みなので、必然的に
	// 複数のGoルーチン間で1つのチャネルを共有すると言うプログラムが現れるってことになる


	ch1 := make (chan int)
	ch2 := make (chan int)

	go reciever(ch1)
	go reciever(ch2)
	// ↑このように並行で走らせることでチャネルを共有しつつ、処理を行うことができる
	// チャネルにデータ入ったらrecieverで出力するって形になるch1


	i := 0

	for i<100{
		ch1 <- i
		ch2 <- i
		time.Sleep(50*time.Millisecond )
		i++
	}




}


