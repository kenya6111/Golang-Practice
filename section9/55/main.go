package main

import (
	"fmt"
	"time"
)

func reciever(c chan int) {
	for {
		i := <-c // chaneelから値を受信する
		fmt.Println(i)
	}

}

// channel
// 複数の後ルーチン間でのデータの受け渡しをするために設計されたデータ構造
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// fmt.Println(<-ch1) //これでチャネルからのデータを受信するとできますが、デッドロックとなる。GOのチャネルは単なるキューではなく、Goルーチン間のデータの共有のために用意されている。
	// このチャネルから受信するという処理は、裏を返せば、他のルーチンがこのチャネルへデータを送信するのを待ちますよって意味。
	// 複数のGOルーチン間でチャネルを共有することが多い

	go reciever(ch1)
	go reciever(ch2)

	i := 0

	for i < 100 {
		ch1 <- i // チャネルに送る
		ch2 <- i // チャネルに送る
		time.Sleep(50 * time.Microsecond)
		i++
	}

}
