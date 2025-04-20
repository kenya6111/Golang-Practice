package main

import (
	"fmt"
	"time"
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
	// ch1 := make(chan int, 2)
	/*
		ch1 <- 1

		close(ch1)  これでチャネルが閉じられる

		// ch1 <- 1

		i, ok := <-ch1 // 第二引数はチャネルが開いているかどうか 厳密にはチャネルの中身が空かつcloseされるとfalseが返される
		fmt.Println(i, ok)
		i2, ok := <-ch1
		fmt.Println(i2, ok)

	*/

	// go reciever("1.goroutin", ch1)
	// go reciever("2.goroutin", ch1)
	// go reciever("3.goroutin", ch1)

	// i := 0
	// for i < 100 {
	// 	ch1 <- i
	// 	i++
	// }

	// close(ch1)
	// time.Sleep(3 * time.Millisecond)
	fmt.Println("----")

	go func() {
		fmt.Println("A start")
		time.Sleep(2 * time.Second)
		fmt.Println("A end")
	}()

	go func() {
		fmt.Println("B start")
		time.Sleep(1 * time.Second)
		fmt.Println("B end")
	}()
	fmt.Println("----")
	time.Sleep(5 * time.Second)
}
