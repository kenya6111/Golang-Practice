package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
// func generator (done chan struct{}, num int) <-chan int{
func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)

	go func() {
		defer wg.Done() // goルーチン終了時にWaitGroupカウントを減らす

	Loop:
		for {
			select {
			// case <-done: // doneチャネルがクローズ次第、return
			case <-ctx.Done(): // doneチャネルがクローズ次第、return
				fmt.Println("ctx.Done is executed")
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
	// done := make(chan struct{})
	// gen := generator(done,1)

	ctx, cancel := context.WithCancel(context.Background())
	gen := generator(ctx, 1)
	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen) // ５回だけ受信

	}

	// close(done)// 5回受信したら、doneチャネルクローズしてキャンセルを実行
	cancel()  // ここでキャンセル通知を送信
	wg.Wait() // generatorの終了を待つ
}
