package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{}{
	// fmt.Println("---")
	// fmt.Println(len(channels))
	// for i:=0; i<len(channels); i++{
	// 	fmt.Printf("%v\n",channels[i])
	// }
	switch len(channels){
		case 0:
			fmt.Println("case0")
			return nil
		case 1:
			fmt.Println("case1")
			fmt.Println(channels[0])
			return channels[0]
	}
	orDone := make(chan interface{})
	// fmt.Printf("ordone: %v\n",orDone)
	// // <- channels[0]
	// // <- channels[1]
	// a:=<- channels[2]
	// fmt.Printf("channels[2]の値: %v\n", a)
	// fmt.Printf("append: %v\n", append(channels[3:], orDone))
	// for i:=range channels{
	// 	fmt.Printf("%v\n",channels[i])
	// }
	go func(){
		defer close(orDone)
		switch len(channels) {
			case 2:
				select{
					case <- channels[0]:
					case <- channels[1]:
				}
			default :
				select {
					case <- channels[0]:
					case <- channels[1]:
					case <- channels[2]:
					case <- or(append(channels[3:], orDone)...):
				}
		}
	}()
	return orDone
}

func signal (after time.Duration) <- chan interface{}{
	done := make(chan interface{})
	go func (){
		defer close(done)
		time.Sleep(after)
	}()
	return done
}

func main() {
	start := time.Now()
	<-or (signal(time.Hour) , signal(time.Minute), signal(time.Second))
	fmt.Printf("done after: %v\n", time.Since(start))
}




// or チャネル

// 複数のキャンセルチャネルがあるときに、
// 「どれか1つでもキャンセルされたら処理を止めたい」ケースがある（例：タイムアウト・ユーザー中断・上位キャンセル）

// <-or(ch1, ch2, ch3)
// ↑これで、ch1・ch2・ch3 のうち1つでも close されたら、or関数からのチャネルが即座に完了するようにしたい。



// 複数のチャネルをまとめて監視して、どれか1つでも完了（＝クローズ）したら即座に反応する	



// case <- or(append(channels[3:], orDone)...)
// の中で呼ばれている or(orDone) は、orDoneチャネルがクローズされない限りブロック状態になる。
// 理由は：「クローズされていないチャネルを受信しようとすると、Goではブロックされる」から。




// case <- channels[2]:
// case <- or(append(channels[3:], orDone)...):

// が別のタイミングで実行されて、

// 先にchannels[2]より先に
// case <- or(append(channels[3:], orDone)...):が実行された場合は、orDoneチャネルが帰ってくるけど受信はできないから、そこでブロック状態になって、

// で、case <- channels[2]:のチャネルがクローズされればこれ呼ばれて、orDoneがクローズされて、

// case <- or(append(channels[3:], orDone)...):の受信が終わって
// mainの中の
// 	<-or (signal(time.Hour) , signal(time.Minute), signal(time.Second))
// のブロックも解除し受信できて、
// やっと
// 	fmt.Printf("done after: %v\n", time.Since(start))が実行できるってわけか



// ✅ あなたがすでに掴んだ重要ポイント
// ポイント	内容
// 受信可能になるには、チャネルに値が送られるか、クローズされる必要がある	クローズされると、<-ch は即座に「ゼロ値 + false（if使うとき）」を返す
// チャネルは基本ブロッキング	誰かが送る／受けるまで止まる（バッファが空／満の場合もブロック）
// select は複数チャネルのうち準備ができたものから選ばれる	順番は**ランダム（公平）**に選ばれるので、実行ごとに違う挙動になる可能性あり
// nil チャネルは常にブロック	誤って nil チャネルを select に入れると永遠にブロックする危険あり