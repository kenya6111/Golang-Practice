package main

import (
	"fmt"
	"math/rand"
	"time"
)



func DoSomething(done chan interface{}) <- chan int{
	readStream := make(chan int)

	go func(){
		defer fmt.Println("DoSomething done")
		defer close(readStream)

		for{
			select{
				case readStream <- rand.Intn(100):
				case <-done:
					return
			}
		}
	}()

	return readStream;
}
func main() {
	done := make(chan interface{})
	readStream := DoSomething(done)

	for i:=1; i<=3; i++{
		fmt.Println(<-readStream)
		// fmt.Println(len(readStream))
	}

	close(done)
	time.Sleep(2*time.Second)
	fmt.Println("main done")
}



// この close(done) が呼ばれた瞬間、done チャネルは クローズ状態 になります。

// 🔄 goroutine 側では：
// for {
// 	select {
// 		case readStream <- rand.Intn(100):
// 		case <-done:
// 			return
// 	}
// }
// done がクローズされると、
// case <-done: が 即座に選ばれるようになる（チャネルからの受信が成功したとみなされる）
// そのため return が実行されて goroutine が終了します。