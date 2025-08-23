package cause

import (
	"context"
	"fmt"
	"time"
)

// ctx が終了した理由を返す
func TestCause() {
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	// <-ctx.Done()
	// fmt.Println("Err:", ctx.Err())            // context deadline exceeded
	// fmt.Println("Cause:", context.Cause(ctx)) // context deadline exceeded

	ctx, cancel := context.WithCancelCause(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel(fmt.Errorf("DB接続エラーで強制キャンセル"))
	}()

	<-ctx.Done()

	fmt.Println("Err:", ctx.Err())            // context canceled
	fmt.Println("Cause:", context.Cause(ctx)) // DB接続エラーで強制キャンセル

}
