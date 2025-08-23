package main

import "fmt"
var data []int = []int{1,2,3,4,5}

func writeToChan(writeChan chan <- int){
	defer close(writeChan)
	for i := range data{
		writeChan <- i
	}

}

func main() {
	handleData := make(chan int)
	go writeToChan(handleData)

	for integer := range handleData{
		fmt.Println(integer)
	}
}

// 拘束→データのアクセスの範囲を縛るもの
// func writeToChan(writeChan chan <- int){ は、int型の値を送信できるチャンネル
// つまり書き込み専用チャネルとして引数で受け取っている
// ✅ writeChan <- 42 ← これは できる
// ❌ <-writeChan ← これはできない（受信は禁止)




// 🧠 何のために制限するの？
// Go の良い設計パターンとして、「関数に余計な権限を与えない」ことがあります。
// 例えば：
// go
// func sendData(ch chan<- int) {
//     ch <- 1 // ✅ OK: 書き込み
//     x := <-ch // ❌ コンパイルエラー（読み込みは禁止）
// }
// これは「この関数は書き込むだけの役割なんだから、受信できちゃダメでしょ？」と 明示するための制限。


// 🔄 逆に読み取り専用にしたいなら？
// func readFromChan(ch <-chan int)
// これは「ch は読み取り専用チャネル」という意味です。
// ✅ x := <-ch はできる
// ❌ ch <- 1 はできない


// 1. 関数の責任が明確になる（副作用を減らす）
// 例：func writeToChan(ch chan<- int) は、「この関数は送信しかしません」と明示している。
// 読み取りの操作をそもそもできなくすることで、関数の役割や責務が明確になる。
// 📌 結果的に、副作用の少ない関数が書ける

// 2. ミスをコンパイル時に検出できる
// 関数が送信用のチャネルを受け取る場合、うっかり x := <-ch みたいに読もうとしたらコンパイルエラーになる。
// 📌 実行時エラーじゃなく、早い段階で防げる = バグが減る。

// 3. 並行設計の意図がより明確に伝わる
// コードを読んでるだけで、「この関数は送るだけ」「この関数は読むだけ」とわかる。
// 特に大規模なプロジェクトやチーム開発で意図がコードに現れるのは非常に重要。
// 📌 コメントよりも信頼性のある「明文化された設計意図」になる。

