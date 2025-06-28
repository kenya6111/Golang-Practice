package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name :=  range []string{"A","B","C"}{
		go func (name string) {
			mutex.Lock()
			defer mutex.Unlock()

			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("Ready ... ")
	time.Sleep(time.Second)
	fmt.Println("Go!")


	cond.Broadcast()

	time.Sleep(time.Second)
	fmt.Println("Done")
}

// Cond
//  条件変数と呼ばれる排他制御の仕組み
//. これもロックをかけたりしてクリティカルセクションを保護するために使われる

// Cond = Condition（条件）
// **NewCond = 「新しい条件変数を作る」**という意味
// コンディション変数は、C言語やJavaにもある古典的な並行処理の仕組み


// cond.Wait() とは？
// 内部的に：
// 	mutex.Unlock()
// 	goroutine を「待機状態」にする
// 	Signal() or Broadcast() が呼ばれたら
// 	mutex.Lock() して Wait() から戻る
// 	→ だから 必ず mutex.Lock() された状態で呼ばないといけない


// cond.wait()で一旦待機（ブロック）
// Signalなどが呼ばれるとcond.Wait()以降の処理が走る


// cond.Broadcast()を使うとSignal()と同様に再開できるが違うってんはwait()を一斉に走らせることができる
