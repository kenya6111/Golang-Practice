package main

import (
	"fmt"
	"sync"
)

var onceA,onceB sync.Once


func A (){
	fmt.Println("A")
	onceA.Do(B)
}
func B (){
	fmt.Println("B")
	onceA.Do(A)
}
func main() {
	// count := 0
	// increment := func(){
	// 	count++
	// }

	// decrement := func(){
	// 	count--
	// }

	// var once sync.Once
	// once.Do(increment)
	// once.Do(increment)
	// once.Do(increment)
	// once.Do(decrement)
	// once.Do(decrement)
	// once.Do(decrement)

	// fmt.Println(count)

	onceA.Do(A)



}


// once.Doで実行すると何回書いても一回しか実行されない

// onceのDoが呼び出されるのが一回という決まりなので、渡す関数が何であろうと、初めに実行され得たDOで終わりになるって感じ

// 上記で言うと最初のonce.Do(increment)しか実行されない


// onceA.Do(A) // → A が呼ばれて
// A() の中で onceA.Do(B) → でも onceA はまだ処理中
// → B() を呼ぼうとしてロック取得しようとしてブロック
// → A() の中で処理が止まる