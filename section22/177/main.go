package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	const timer = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()
		count :=0

		begin := time.Now()

		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(3 * time.Nanosecond)
			lock.Unlock()
			count++
		}

		fmt.Printf("greedyWorke: %v\n", count)
	}

	politeWorker := func(){
		defer wg.Done()

		count :=0

		begin := time.Now()

		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			count++
		}

		fmt.Printf("politeWorker: %v\n", count)
	}

	wg.Add(2)

	go greedyWorker()
	go politeWorker()

	wg.Wait()


	// test
	// var count2 int
	// wg.Add(100)
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		lock.Lock()        // 🔐 ロック
	// 		count2++          // 🟢 安全な操作
	// 		lock.Unlock()      // 🔓 アンロック
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("安全なカウント:", count2)
	// Lock() で「今は自分だけが count を触っていい」を保証


	// test
	// var count2 int
	// wg.Add(100)
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		count2++          // ❗競合状態（安全じゃない）
	// 		defer wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("安全なカウント:", count2)
	// ✅ 実行してみよう：
	// たいてい 100 にならない。毎回違う数字になる（競合状態）。


}

// リソース枯渇
// greedyWorker	ロックを1回長めに取ってから作業 → 繰り返す
// politeWorker	ロックを短く3回に分けて細かく使う → 繰り返す
// 	どちらも「1秒間で何回処理できるか（= count）」を競っている
// 	処理のたびに lock.Lock() して Sleep() して Unlock() を繰り返している

// 💥 ポイント：「リソースの使い方の姿勢（greedy/polite）」で性能が変わる
// 	🔴 greedyWorker：
// 	1回ロックを取ったら長めに保持
// 	他のゴルーチンがロックを使いたくても「空くまで待たされる」
// 	→ 他の処理を“邪魔”しやすい（貪欲＝greedy）

// 	🟢 politeWorker：
// 	ロックを短く、こまめに分けて使う
// 	毎回ロック・解除を丁寧にやることで、他のゴルーチンにもチャンスを与える
// 	→ より「協調的（polite）」に動く


// greedyWorker 500000
// politeWorker 100000
// 	のように、greedyWorker の方が高い数字を出すことがあります。
// 	でも、これは「うまくやった」というより：
// 	「他人がロックを使おうとしてるのを横取りして、一人でずっと使い続けた」状態です。


// greedyWorker がロックを握りすぎることで、politeWorker が リソースを使えなくなる
// politeWorker はチャンスを伺ってるのに、ずっと待たされる
// これは一種の「リソース飢餓状態（starvation）」


// Mutex = Mutual Exclusion
// （ミューチュアル・エクスクルージョン）
	// Mutual（相互）：お互いに
	// Exclusion（排除）：除外する
	// つまり：
	// 「お互いに排除し合う」＝ 同時に処理させないようにする仕組み
	// “A mutex is used to ensure mutual exclusion.”
	// “ミューテックスは 相互に排除しながら 安全にデータにアクセスするための仕組みです。”