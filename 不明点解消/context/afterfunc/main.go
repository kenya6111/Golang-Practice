package afterfunc

import (
	"context"
	"fmt"
	"time"
)

func TestAfterFunc() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	stop := context.AfterFunc(ctx, func() {
		fmt.Println("context was canceled")
	})

	time.Sleep(2 * time.Second)
	// fmt.Println("stop canceled ", stop()) // stopは、キャンセルされたときに実行する予定の関数を、もしまだ実行されていなければキャンセルする
	// cancel()
	cancel()
	time.Sleep(2 * time.Second)

	if stop() {
		fmt.Println("後処理をキャンセルできた！")
	}
	// time.Sleep(500 * time.Millisecond)

}
