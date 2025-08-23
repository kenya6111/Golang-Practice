package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
// func generator (done chan struct{}, num int) <-chan int{
func generator (ctx context.Context, num int) <-chan int{
	out:= make(chan int)

	go func(){
		defer wg.Done()// goルーチン終了時にWaitGroupカウントを減らす
	
		Loop:
			for{
				select{
				// case <-done: // doneチャネルがクローズ次第、return
				 case <-ctx.Done(): // doneチャネルがクローズ次第、return
					break Loop
				case out <- num:// キャンセルされるまでずっとnumを送信
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
	
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	gen := generator(ctx,1)
	wg.Add(1)

	LOOP:
	for i:= 0; i<9000000000000000000;i++{
		select{
			case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}

	}

	// close(done)// 5回受信したら、doneチャネルクローズしてキャンセルを実行
	cancel()// ここでキャンセル通知を送信
	wg.Wait()// generatorの終了を待つ
}

