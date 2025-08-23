package main

import "fmt"



func DoSomething(strings <- chan string) <- chan interface{}{
	completed := make(chan interface{})
	fmt.Println(1)
	
	go func (){
		defer fmt.Println("DoSomething Done")
		defer close(completed)

		fmt.Println(2)
		for s := range strings {
			fmt.Println(s)
		}
		fmt.Println(5)
	}()
	fmt.Println(3)
	return completed
}
func main() {
	completed := DoSomething(nil)// nilチャネルを渡す
	<-completed
	fmt.Println("main Done")
}

// Goルーチンのキャンセル処理
// プロセス内に残ったゴルーちんは、がベジコレクションでメモリが解放されない
//

// ❗ range nilチャネル は何が起きるか？
// for s := range strings の strings が nil だと、range は永遠に止まる
// これは 受信が永遠に来ない状態＝ブロック


// 🧠 補足：nil チャネルの性質
// 送信 (ch <- x) → 永遠にブロック
// 受信 (<- ch) → 永遠にブロック
// range ch → 永遠にブロック
// select { case <-ch: } → 他の case がないとブロック
// つまり：
// nil チャネルは完全なストッパー。使い方を誤ると確実に deadlock になる



// 🔥 結論：range nil チャネル は 永遠にブロックされる
// DoSomething(nil) として nil チャネルを渡す
// goroutine の中で for s := range strings が実行される
// しかし strings は nil
// → range nil は データが来ないままずっと待機（ブロック