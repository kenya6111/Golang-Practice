package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
				// case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")

	}()
	return out
}
func main() {
	done := make(chan struct{})
	gen := generator(done, 1)
	deadlineChan := time.After(3 * time.Second) // 指定した時間が経過すると、値を1つだけ送ってくるチャネル（<-chan time.Time）を返す関数です。

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result := <-gen:
			fmt.Printf("result:%v\n", result)
		case <-deadlineChan:
			fmt.Println("timeout")
			break LOOP
		}
	}
	close(done)
	wg.Wait()
}

// time.After(...) は 内部で自動的に1回だけ値を送ってくれるチャネルです。
// deadlineChan := time.After(1 * time.Second)
// これをすると：
// 「1秒後に、現在時刻（time.Time型の値）を1回だけ送信してくるチャネル」が返ってくる
