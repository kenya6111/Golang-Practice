package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	cond := sync.NewCond(&sync.Mutex{})

	go func (){
		for range time.Tick(1 * time.Second){ // 1秒おきに実行する	
			cond.Broadcast()
		}
	}()

	var flag [2]bool

	takeStep := func (){
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
	}

	var wg sync.WaitGroup

	p0 := func(){
		defer wg.Done()
		flag[0] = true
		takeStep()
		for flag[1] {
			takeStep()
			flag[0] = false
			takeStep()
			if flag[0] != flag[1] {
				break
			}
			takeStep()
			flag[0] = true
			takeStep()
			fmt.Println("p0")
		}
	}
	p1 := func(){
		defer wg.Done()
		flag[1] = true
		takeStep()
		for flag[0] {
			takeStep()
			flag[1] = false
			takeStep()
			if flag[0] != flag[1] {
				break
			}
			takeStep()
			flag[1] = true
			takeStep()
			fmt.Println("p1")
		}
	}

	wg.Add(2)

	go p0()
	go p1()

	wg.Wait()
}

// ✅ ライブロック（Livelock）とは？
// お互いに譲り合いすぎて、結果的にどの処理も前に進めなくなる状態
// （止まってるわけではない、でも進んでない）
// お互いにブッロクはされてない

// goroutine やスレッドが「他の処理のために一時停止したり再試行したり」を繰り返す
// けど、何回譲ってもタイミングが悪くて誰も前に進めない
// 動いてはいるけど、本質的には“永久に足踏みしてる”状態

// ✅ 例えるなら（実生活）
// 2人が細い通路で鉢合わせ
// お互い譲ろうとして、左によけて→ぶつかって→右によけて→またぶつかって…
// → 譲り合ってるのに一生進まない



// ✅ 語源：Live + Lock
// 	単語	意味
// 	Live	生きてる、動いてる（≠停止）
// 	Lock	ロック状態にある、進めない

// 	🔤 つまり：
// “動いてる（live）のに進まない（lock）”というパラドックスな状態

// ラブロックはエラーを吐かないのでデッドロックより見つけるのが難しい。
