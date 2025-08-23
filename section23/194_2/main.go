package main

import "fmt"
var data []int = []int{1,2,3,4,5}


func chanOwner () <-chan int {
	result := make(chan int, 5)
	go func(){
		defer close(result)

		for i:=1; i<=50; i++{
			result <- i
		}
	}()

	return result
}

func consumer (result <- chan int ){
	for result := range result{
		fmt.Println(result)
	}
}
func main() {
	result := chanOwner()
	consumer(result)

}


// チャネルの拘束化によるメリット
// 1. チャネルの所有権（責任の所在）を明確にできる
// func chanOwner() <-chan int
// chanOwner はチャネルを 生成・送信だけ する関数。
// 返すのは 読み取り専用チャネル（<-chan int）だから、呼び出し元は「読むことしかできない」＝チャネルの送り手側を勝手にいじれない。
// 「チャネルを作った関数が、その送信の責任も持つ」という明確な設計意図を保証する。

// 2. 受信側（consumer）は「読み取りしかできない」ことが保証される
// func consumer(result <-chan int)
// これは「受け取るだけの関数」だと明示。
// result <- 1 のような書き込み操作はコンパイルエラーになるので、意図しない変更を防止できる。
// 読み取り専用としてコードレベルで封じ込めるので、安全。


// 3. 誤操作を防げる（安全性アップ）
// たとえば、うっかり以下のようなコードを書いたとしても：
// result <- 999 // 書き込みのつもり
// 型制約があるため、コンパイル時にエラーになる：
// cannot send to receive-only channel
// これにより、開発者のミスを未然に防止できる。


// 4. コードの意図が明確になり、読みやすくなる
// func consumer(result <-chan int)
// と見れば、「あ、この関数は result を読むだけなんだな」とドキュメント不要で分かる。
// また chanOwner() の返り値も <-chan int なので、「これは読み取り専用としてしか使わない前提のチャネルだな」と明確になる。


// 結論：拘束は「制限」じゃなく「設計ミス防止と責任分離」
// 方向を制限することで：
// チャネルの「送り手」と「受け手」の役割を明確化
// 誤操作やバグをコンパイル時に検出
// チーム開発での意図伝達が明快
// Goでは 「明示的に制限することが安全で正しい設計」 とされており、これはまさにその好例です。

