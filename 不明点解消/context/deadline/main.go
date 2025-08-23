package deadline

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// キャンセルされるまでnumをひたすら送信し続けるチャネルを生成
func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			fmt.Println(11)
			select {
			case <-done: // doneチャネルがcloseされたらbreakが実行される
				fmt.Println(22)
				break LOOP
				// case out <- num: これが時間がかかっているという想定
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func TestDeadline() {
	done := make(chan struct{})
	gen := generator(done, 1)
	fmt.Println(33)
	deadlineChan := time.After(time.Second * 4)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		fmt.Println(55)

		select {
		case result := <-gen: // genから値を受信できた場合
			fmt.Println(result)
		case <-deadlineChan: // 1秒間受信できなかったらタイムアウト
			fmt.Println("timeout")
			break LOOP
		}
	}
	close(done)

	wg.Wait()
}
func generator2(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			fmt.Println(11)
			select {
			case <-ctx.Done(): // doneチャネルがcloseされたらbreakが実行される
				fmt.Println(22)
				break LOOP
				// case out <- num: これが時間がかかっているという想定
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func TestDeadline2() {
	// done := make(chan struct{})
	// gen := generator2(done, 1)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))
	gen := generator2(ctx, 1)
	fmt.Println(33)
	// deadlineChan := time.After(time.Second * 4)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		fmt.Println(55)

		select {
		// case result := <-gen: // genから値を受信できた場合
		// 	fmt.Println(result)
		// case <-deadlineChan: // 1秒間受信できなかったらタイムアウト
		// 	fmt.Println("timeout")
		// 	break LOOP
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}
	// close(done)
	cancel()

	wg.Wait()
}
