package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	value int
	name string
}
func main() {
	var wg sync.WaitGroup
	printSum := func(v1,v2 *value){
		defer wg.Done()
		v1.mu.Lock()
		fmt.Printf("%vがロックを取得しました\n", v1.name)

		defer v1.mu.Unlock()

		time.Sleep(2*time.Second)

		v2.mu.Lock()
		fmt.Printf("%vがロックを取得しまいした\n", v2.name)

		defer v2.mu.Unlock()

		fmt.Println(v1.value + v2.value)
	}

	var a value = value{
		name:"a",
	}
	var b value = value{
		name:"b",
	}

	wg.Add(2)

	go printSum(&a,&b)
	go printSum(&b,&a)

	wg.Wait()
}


// デッドロック
// お互いが相手の処理を終わるのを待っていて、処理終わらないこと

// 上記の処理はデットロックが発生してしまう。
// なぜか。
// 	go printSum(&a,&b)が「 A 」、	go printSum(&b,&a)が「 B 」とすると、
//   A
// Goroutine A: ロック(a) → スリープ → ロック(b) ← ✖ bはBが保持してるからブロックされる
// Goroutine B: ロック(b) → スリープ → ロック(a) ← ✖ aはAが保持してるからブロックされる
// って感じになり、デッドロックする

// これは、printSumのUnlockがdefer付きなので、aとbをロックして、printSumの最後まで行ったら、両方アンロックする
// って処理の構造になっているから。

// →ようは2つ目のロックも取らないとアンロックできない構造なのでデッドロックが発生している
// 厳密というかそこまで厳密ではないが　→　「次のロックを取らないと、ロック解除しない」わけではない。。。

// v1.mu.Lock()
// defer v1.mu.Unlock()
// time.Sleep(2 * time.Second) // ← この間ずっとv1はロックされたまま
// v2.mu.Lock()
// defer v2.mu.Unlock()

// となっているので、2つ目のロック取らないと2つともロック解除始まらない順番になっている。だけ。
