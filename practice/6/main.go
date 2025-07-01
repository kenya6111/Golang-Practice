package main

import "fmt"


func main(){
 	ch1 := make(chan int, 2)// 初期化時はチャネルはオープン状態
	ch1 <- 1
	ch1 <- 2
	close(ch1)

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	v,ok := <-ch1
	fmt.Println(v,ok)	// チャネルが空でクローズ済み：ゼロ値 + false が返る


	// fmt.Println(<-ch1)


	// i,ok := <-ch1
	// fmt.Println(i, ok)// ここのokがチャネルがオープンがクローズかどうか


}


