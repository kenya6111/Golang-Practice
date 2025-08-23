package main

import (
	"fmt"
	"log"
	"os"
)

type Result struct {
	Response *os.File
	Error error
}
func CheckFiles (done <-chan interface{}, filenames ...string) <-chan Result{
	results := make(chan Result)
	go func(){
		defer close(results)

		for _,filename := range filenames {
			var result Result

			file, err := os.Open(filename)
			result = Result{file,err}

			select{
				case <-done:
					return
				case results <- result:
			}
		}

	}()
	return results

}

func main() {
	done := make(chan interface{})

	defer close(done)

	filenames:= []string{"main.go","x.go"}

	for result:= range CheckFiles(done, filenames...){
		if result.Error != nil{
			log.Printf("error: %v\n", result.Error)
			break
		}

		fmt.Printf("Response: %v\n", result.Response.Name())
	}

}

// エラーハンドリングについての考え方
// エラーハンドリング誰がそのエラーを処理する責任を負うのかを考えることが大事

// 結論から言うと 「チャネルが閉じられる」ことも、受信できる条件になります。

// 🔁 Go のチャネル受信の挙動まとめ
// val, ok := <-ch
// この時の動作はこう
// 状況	                                      valの値     	okの値	 備考
// チャネルに値あり	                             その値	       true	   普通の受信
// チャネルが閉じられていて、中に値がまだ残っている	   その値      	true	残ってるぶんは読める
// チャネルが閉じられていて、値も空っぽ	             ゼロ値	      false	  完全にクローズ＆空 → ok == false

// case <-done: の意味
// この select の中の case <-done: は、次のような意味になります：

// select {
//   case <-done:
//     // ここに来たら done チャネルが閉じられていた or 値が送られてきた
//     return
//   case results <- result:
//     // done が来てなければ通常通り送信する
// }
// つまり：
// done に 誰かが値を送った
// または、done が クローズ（close(done)）された
// このどちらでも発火します！