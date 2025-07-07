package main

import "fmt"

func main(){
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)


	go func (){
		for{
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func (){
		for {
			i2 := <-ch4
			ch5 <- i2 -1
		}
	}()


	n:=0

	for{
		select{
			case ch3 <- n:
					n++
			case i3 := <- ch5 :
				fmt.Println("recieved", i3)
		}

		if n > 5 {
			break
		}
	}

}


// ⏱ 例えば n = 0 のときの処理の流れ
// main の select で ch3 <- 0 が成功（空いてるので送信できる）
// goroutine① が ch3 から 0 を受信し、0 * 2 = 0 を ch4 へ送信
// goroutine② が ch4 から 0 を受信し、0 - 1 = -1 を ch5 へ送信
// 次に select の case i3 := <-ch5 が実行され、main が -1 を受け取る
// received -1 が出力される

// ⏱ n = 1 のときも同じ：
// main で ch3 <- 1 が送信される
// goroutine①： 1 を受信して 2 を ch4 に送る
// goroutine②： 2 を受信して 1 を ch5 に送る
// main で 1 を受信して出力 received 1

