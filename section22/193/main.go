package main

import (
	"fmt"
	"time"
)

func main() {
	// a := make(chan int)
	// b := make(chan int)

	// close(a)

	// select{
	// 	case <-b:
	// 	case <-a:
	// }

	start := time.Now()

	ch1 := make(chan int)
	ch2 := make(chan int)

	done := make(chan interface{})

	go func(){
		time.Sleep(2 * time.Second)
		close(done)
	}()



	go func (){
		defer close(ch1)
		for i := 0; i< 10; i++{
			time.Sleep(1 * time.Second)
			ch1 <-i
		}
	}()
	go func (){
		defer close(ch2)
		for i := 0; i< 10; i++{
			time.Sleep(1 * time.Second)
			ch2 <-i
		}
	}()

	loop:
		for{
			select{
				case <- done:
					break loop
				case v, ok := <-ch1:
					if !ok{
						break loop
					}
					fmt.Printf("ch1: %v\n",v)
				case v, ok := <-ch2:
					if !ok{
						break loop
					}
					fmt.Printf("c21: %v\n",v)
			}
		}
		end := time.Now()

		fmt.Println(end.Sub(start))

}

// チャネル→複数のごルーチンを結ぶものだったが

// select →複数のチャネルの操作をまとめるもの


// select は Goの並行処理構文の一つで、「複数のチャネル操作からどれか一つを選択して実行する」ためのものです。
// 今、複数のチャネルを監視してるけど、先に届いたやつだけ処理するよ」
// →「select（選ぶ）」ってこと。