package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func(){
		time.Sleep(2*time.Second)
		ch <- "hello"
	}()

	fmt.Println(<-ch)
}


// channel GOルーチン間のメモリに対するアクセス同期と、GOルーチン間の通信に使われる
// チャネルはキューというfirst in first out型の

// ちゃねるはFIFOなのでデータを順序よく受け渡すデータ構造
// 並行処理で正しくデータを受け渡す同期機構
// 	ちゃねるは生合成が壊れる心配がないので安全。なのでGOルーチン間の通信にチャネルをよく使う
// チャネルは読み書きが準備できるまでブロックします


// 上記のソースではWaitGroupを使っていないので、time.Sleepがある時点で先にfmt.Println()の出力まで行ってしまって
// 何も出力されず終わるのでは？？と思うかもしれないが、
// チャネルは入力が来るまで読み込みは待つので、上記でもしっかり入力が来るまで、fmt.Println(<-ch)の実行はされないっていうことが起こる
	// チャネルがブロックしてくれるってわけだね