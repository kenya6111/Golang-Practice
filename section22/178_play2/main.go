package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rw sync.Mutex
	var wg sync.WaitGroup

	read := func(id int) {
		defer wg.Done()
		rw.Lock()
		defer rw.Unlock()

		fmt.Printf("👀 Reader %d: start\n", id)
		time.Sleep(1 * time.Second)
		fmt.Printf("👀 Reader %d: end\n", id)
	}

	wg.Add(3)
	go read(1)
	go read(2)
	go read(3)

	wg.Wait()
}

// Mutexを使った場合の処理
// k_tanaka@tanakakenyanoMacBook-Air 178_play2 % go run main.go 
// 👀 Reader 3: start
// 👀 Reader 3: end
// 👀 Reader 1: start
// 👀 Reader 1: end
// 👀 Reader 2: start
// 👀 Reader 2: end

// RWMutexを使った場合の処理
// k_tanaka@tanakakenyanoMacBook-Air 178_play2 % go run main.go 
// 👀 Reader 3: start
// 👀 Reader 2: start
// 👀 Reader 1: start
// 👀 Reader 1: end
// 👀 Reader 2: end
// 👀 Reader 3: end
// k_tanaka@tanakakenyanoMacBook-Air 178_play2 % 


// ✅ 結論：
// RLocker()（＝RLock()）を使った処理は、他の RLock() / RLocker() を使った処理と
// 同時に並行実行できます。お互いにブロックしません！


// 🧍‍♂️ Goroutine A: RLock() → 進行中...
// 🧍‍♀️ Goroutine B: RLock() → 同時に入れる！
// 🧍 Goroutine C: RLock() → これもOK！
// ⬅︎ 読み取り処理は並行して入ってこれる！


// RLock() は「今読み取り中なんで、書き込み入れないでね」って宣言してるだけ。
// 読み取り同士ではブロックし合わないけど、
// 「書き込みから読み取りを守っている」ことが本質なんです。


// RLocker() の処理中、他の RLocker() の処理は入れる？	✅ YES！
// 書き込み（Lock()）が来たらどうなる？	⏸ RLock() たちが終わるまで 待機する