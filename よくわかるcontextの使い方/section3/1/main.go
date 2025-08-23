package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()

	Loop:
		for {
			select {
			case <-done: // doneチャネルがクローズ次第、return
				break Loop
			case out <- num: // キャンセルされるまでずっとnumを送信
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}
func main() {
	done := make(chan struct{})
	gen := generator(done, 9)
	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)

	}

	close(done) // 5回受信したら、doneチャネルクローズしてキャンセルを実行

	wg.Wait()
}
