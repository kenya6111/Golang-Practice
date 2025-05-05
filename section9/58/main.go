package main

import (
	"fmt"
)

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")

}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	ch2 <- "A"
	ch1 <- 1
	ch1 <- 1
	ch2 <- "B"

	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000) // チャネル1にデータが入っていた場合はこっちが実行
	case v2 := <-ch2: // ちなみに上のcaseが優先されるわけではなく、実行できるcaseが複数ある場合は、どのケースを実行できるかは「ランダム！！」
		fmt.Println(v2 + "!?") // チャネル2にデータが入っていた場合はこっちが実行

	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// reciever
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 4
		}
	}()

	//
	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
		}

		if n > 100 {
			break
		}
	}

}

// | 項目            | `switch`        | `select`                          |
// | ------------- | --------------- | --------------------------------- |
// | 主な用途          | 値に応じて条件分岐する     | チャネルの通信（送受信）を監視して分岐               |
// | 対象            | 値や条件式           | チャネルの送信 or 受信                     |
// | 同時に成立         | 基本的に1つ（上から順に評価） | 複数 case が ready な場合は**ランダム**に選ばれる |
// | `default` の意味 | 条件にマッチしないときに実行  | すべてのチャネルがブロック中のときに即実行             |
// | 用途例           | 値による処理分岐        | goroutine 間の非同期通信制御など             |

// select の具体的な特徴
// select {
// case v := <-ch1:
//     // ch1 から受信できたら実行
// case ch2 <- 100:
//     // ch2 に送信できたら実行
// default:
//     // すべてブロック中なら即実行（非ブロッキングにできる）
// }
// 複数のチャネル操作を同時に待てる

// どれか一つでも ready ならランダムで実行される

// 非同期プログラミングやタイムアウト処理にめちゃ便利
