package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var memoryAccess sync.Mutex
	var data int


	wg.Add(1)

	go func(){
		defer wg.Done()
		memoryAccess.Lock()
		data ++
		memoryAccess.Unlock()
	}()

	wg.Wait()

	// time.Sleep(1 * time.Second)
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	}else{
		fmt.Println(data)
	}
	memoryAccess.Unlock()


	// var count int

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	count++ // A
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	count++ // B
	// }()
	
	// wg.Wait()
	// fmt.Println(count)
}

// 競合状態
// 順番が保証されていない状態が競合状態
// 出力結果にばらつきが出てしまいバグの原因になる
// メモリのアクセス同期
// クリティカルセクション
// sync mutex →アクセス同期
// wgは後ルーチンが終わるまで待ちましょうのやつ
// 複数のGo routineが同じデータにアクセスするのを防ぐやつ
// wgだけではデータアクセス競合は防げない。終了を同期するだけ
// なのでMutexは共有変数に複数の語ルーチンからアクセスするときに必要になる

// クリティカルセクション⇩
//  同時に複数のスレッドやgoroutineが実行してはいけない、共有資源へのアクセス部分のこと。

// ✅ 1. sync.WaitGroup（＝終了を待つ）
// 🎯 目的：
// 複数の goroutine の終了を待つために使う
// 📦 機能：
// Add(n)：待ちたいゴルーチンの数を増やす
// Done()：ゴルーチン内で「終わったよ」と通知
// Wait()：すべての Done() が呼ばれるまで待機



// ✅ 2. sync.Mutex（＝同時アクセス防止）
// 🎯 目的：
// 同じ変数への同時アクセス（競合状態）を防ぐために使う
// 📦 機能：
// Lock()：排他ロック開始（他の goroutine はここでブロックされる）
// Unlock()：ロックを解除（次の人が通れる）



// 競合状態とは、複数の処理が同じ変数にどちらかの処理が書き込む前にアクセスしてしまい、同じ値を読み込んでしまい、結果がおかしくなること
// 下記の処理だと、A処理が書き込む前にBもおを読み取ってしまい、結果どちらの処理も０→１にカウントアップしてしまっている。
// 📌 たとえばこんなケース：
// var count int

// go func() {
//     count++ // A
// }()

// go func() {
//     count++ // B
// }()
// ここで count++ が 2 回実行されるはずだけど……

// 処理の順番次第では count = 1 にしかならないことがある

// これは、count++ が実は「読み込み → 加算 → 書き込み」の3ステップで構成されているから

// 😱 問題の原因
// 複数のgoroutineが同時に「読み込み」→「書き込み」してしまうと：

// goroutine A	 goroutine B	 結果
// 読み込み: 0	   読み込み: 0	    どっちも 0 を読んでる
// 加算: 1	      加算: 1	       両方とも 1 になると思ってる
// 書き込み: 1	   書き込み: 1	     最終的に count = 1 ← ❌





