package donetest

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done: // doneチャネルがcloseされたらbreakが実行される
				break LOOP
			case out <- num: // キャンセルされてなければnumを送信
			}
			num++
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func TestDone() {
	done := make(chan struct{})
	gen := generator(done, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done) // 5回genを使ったら、doneチャネルをcloseしてキャンセルを実行

	wg.Wait()
}
func generator2(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num: // キャンセルされてなければnumを送信
			}
			num++
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func TestDone2() {
	// done := make(chan struct{})
	// gen := generator2(done, 1)
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator2(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel() // 5回genを使ったら、doneチャネルをcloseしてキャンセルを実行

	wg.Wait()
}
