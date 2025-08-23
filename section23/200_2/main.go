package main

import "fmt"

func generator (done <-chan interface{}, integer ...int) <-chan int{
	intStream := make(chan int)

	go func(){
		defer close(intStream)

		for _,v := range integer{
			select{
			case <-done:
				return
			case intStream <-v:
			}
		}
	}()

	return intStream
}

func double (done <-chan interface{}, intStream <-chan int) <-chan int{
	doubleStream := make(chan int)

	go func (){
		defer close(doubleStream)
		for i := range intStream {
			select{
				case <-done:
					return
				case doubleStream <- i*2:
			}
		}
	}()

	return doubleStream
}

func add (done <-chan interface{}, intStream <-chan int) <-chan int{
	addStream := make(chan int)

	go func (){
		defer close(addStream)
		for i := range intStream {
			select{
				case <-done:
					return
				case addStream <- i + 1:
			}
		}
	}()

	return addStream
}


func main() {
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1,2,3,4,5)
	for v :=range double(done, add(done, intStream)) {
		fmt.Println(v)
	}
}


// パイプライン


// generator
// done には何も送られていない（または閉じられてもいない）→ case <-done: はブロック
// intStream <- v の方は送信先がまだ生きていて受信される余地がある → 選ばれるのはこっち
// したがって、実行中は done 側の case は 「選ばれていない」（＝呼ばれていない）ということになります。


// 🧠 なんで「パイプライン」っていうの？
// 工場のパイプラインと同じ意味です。
// 工場では材料が「流れて」、途中で加工されたり検査されたりして、最終的に製品になりますよね。Goでも同様に、「値（データ）」が流れていき、各段階で加工されていきます。


// 🛠 具体例（このコード）
// intStream := generator(done, 1, 2, 3, 4, 5)
// for v := range double(done, add(done, intStream)) {
//     fmt.Println(v)
// }
// これはこう流れます：
// generator: 値を生成 → チャネルに送信
// ⇩
// add: 受け取った値に +1 して次へ送信
// ⇩
// double: 受け取った値を *2 して次へ送信
// ⇩
// main: 最終結果を range で出力



// main: for v := range double(...) {
//         4     → OK
//         6     → OK
//         8     → v > 6 → break
//      } ↓ループを抜ける
// ↓ mainの defer close(done) が呼ばれる
// ↓ 全goroutine内の select { case <-done: } が発火
// ↓ 各 goroutine は return（終了）
// ✓ 安全に全処理が終了！

// この設計の本質
// done は「ゴルーチンにやめてくれ！」と伝える非常停止ボタン
// break だけでは止まらない
// done を閉じると全体が安全に終了する



// Goのチャネルの設計はこうなっています：
// チャネルは 「ストリーム」＝逐次データの流れ
// range ch はチャネルに値が届くたびに1つずつ受け取る
// チャネルが close() されたときにようやく range が終了する



// 🔄 動作の流れ（ステップで）
// double(...) が呼ばれる
// make(chan int) でチャネル作成
// go func(...) がスタート（まだ何も送ってない）
// 即座に return doubleStream でチャネル本体だけ返す
// main の range double(...) が始まる（まだ値は来てないので待つ）
// goroutine 内部で値が1つ送られる
// range の1ループ目が回る
// 値が送られ続ける間、ループが続く
// 全部送ったら close(doubleStream) され、range も終了




// generator
// intStream というチャネルを作る
// 新しい goroutine を起動する
// 1,2,3,4,5 を for で順に送っていく
// generator() 関数自体は、チャネル intStream をすぐ返す
// （値はまだ何も送ってない可能性がある）
// この時点で、裏でデータを順に送信する goroutine が走り始めている

// add
// intStream から値を1個受け取るごとに i + 1 して addStream に送る
// goroutine で実行されていて、処理しながらデータを送っていく
// チャネル addStream をすぐ返す（値はまだない）

// double
// addStream から1つずつ受け取り、i * 2 して doubleStream に送信
// goroutine で実行
// チャネル doubleStream を返す（値はまだ届いていない）


// main の range ループ開始
// for v := range double(...) {
// 	fmt.Println(v)
// }
// ここで doubleStream を range する
// これは「チャネルから1個ずつ値が届くたびにループを1回回す」という意味
// 1個目の値が届くまでブロックして待つ

// ✅ 7. 処理はどの時点で終わるのか？
// generator が intStream に全ての値を送り終えたら close(intStream)
// add の range intStream が終了 → close(addStream)
// double の range addStream が終了 → close(doubleStream)
// main の range doubleStream が終了 → main() 完了 → defer close(done)