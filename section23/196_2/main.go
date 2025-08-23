package main

import (
	"fmt"
	"time"
)



func DoSomething(done<-chan interface{}, strings <- chan string) <- chan interface{}{
	completed := make(chan interface{})

	go func (){
		defer fmt.Println("DoSomething Done")
		defer close(completed)

		for{
			select{
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
			}
		}
	}()

	return completed
}
func main() {
	done := make(chan interface{})
	completed := DoSomething(done,nil)

	go func(){
		time.Sleep(2*time.Second)
		close(done)
	}()
	<-completed
	fmt.Println("main Done")
}


// プロセス内に残ったゴルーちん
// Loop
//チャネル
// close
// go goroutine
// select とは


//キャンセル処理



// for {
// 	select {
// 		case s := <-strings:
// 			fmt.Println(s)
// 		case <-done:
// 			return
// 	}
// }
// select は どれか1つの case が準備できたら実行
// strings は nil なのでブロックされっぱなし（スキップ）
// 2秒後に done が close され、case <-done が実行される
// → return して goroutine 終
// → defer close(completed) で完了通知

