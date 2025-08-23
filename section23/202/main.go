package main

import (
	"fmt"
	"math/rand"
	"time"
)

func repeatFunc (done <-chan interface{}, fn func() interface{}) <-chan interface{}{
	valueStream :=make(chan interface{})
	go func(){
		defer close(valueStream)
		for{
			select{
			case <- done:
				return 
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{}{
	takeStream:= make(chan interface{})

	go func(){
		defer close(takeStream)

		for i:=0; i<num; i++ {
			select{
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int{
	intStream:= make(chan int)

	go func(){
		defer close(intStream)

		for v:=range valueStream {
			select{
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()

	return intStream
}
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{}{
	primeStream:= make(chan interface{})

	go func(){
		defer close(primeStream)

		L:
		for i:=range intStream {
			for div := 2; div<i; div++{
				if i%div==0{
					continue L
				}
			}

			select {
			case <-done:
				return
			case primeStream <-i:
			}
		}
	}()

	return primeStream
}

func random() interface{}{
	return rand.Intn(5000000000)

}
func main() {
	done := make(chan interface{})

	defer close(done)

	randIntStream:= toInt(done, repeatFunc(done,random))

	start :=time.Now()

	for prime := range take(done,primeFinder(done, randIntStream), 10){
		fmt.Println(prime)
	}

	fmt.Println(time.Since(start))
}

// 上記でだいたい20秒くらいかかる
// これをファンイン、ファンアウトを使って効率化できる

// ファンイン
// ファンアウト

// パイプラインの特定のステージで計算量が多くなってしまいパイプライン全体に影響及ぼして遅くなること
// ファンアウト（fan-out）	1つの入力から、複数の並行処理に仕事をばらまくこと
// ファンイン（fan-in）	複数の処理結果を1つの出力に集約すること

// fan = 扇
// fan-out = 扇状に広げる（= 分散）
// fan-in = 扇状に閉じる・集める（= 集約）

// 🌀 ファンアウト
// 1つのデータ → 複数のgoroutineへ
//            +----→ worker1
// source --->+----→ worker2
//            +----→ worker3

// 🔁 ファンイン
// 複数のデータ → 1つにまとめる
// worker1 ─┐
// worker2 ─┼─→ mergedOutput
// worker3 ─┘

// ファンアウト	重い処理を分散して並列で速くする（例：画像変換、API呼び出し）
// ファンイン	並列で処理した結果をまとめて使うため（例：統合、出力、DB保存）