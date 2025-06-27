package main

import (
	"fmt"
	"sync"
	"time"
)


func sayHello (wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hello world")
}

func runTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("タスク %d 開始\n", id)
	time.Sleep(time.Duration(id) * time.Second) // 擬似的な処理
	fmt.Printf("タスク %d 終了\n", id)
}


func main() {
	var wg sync.WaitGroup
	// wg.Add(1)

	// go sayHello(&wg)

	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("hello")
	// }()

	// wg.Wait()

	for  i:=1; i<=5; i++{
		wg.Add(1)
		go runTask(i, &wg)
	}
	wg.Wait()
	fmt.Println("すべてのタスクが完了しました")
}

// goキーワードを使うことで関数を並行して時効できる
// go sayHello などで別スレッドで非同期実行できる
// sync.WaitGroup → Goroutineの終了を待つための機構
// wg.Add(n) n個のgoroutineを待つようにカウントを増やす
// 実行中のgoroutineが終わったときに呼ぶ。Addに対応してカウントを-1する働きがある
// 全てのカウンと（Add()した個数）が０になるまでブロックする

//  func (){
//  }()で即実行できる無名関数

