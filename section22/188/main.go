package main

import (
	"fmt"
	"sync"
)

func main() {
	// ch := make(chan int)

	// go func (){
	// 	// defer close(ch)

	// 	for i:=0; i<10; i++{
	// 		ch <- i
	// 	}
	// }()

	// for integer := range ch{
	// 	fmt.Println(integer)
	// }

	begin := make(chan interface{})

	var wg sync.WaitGroup

	for i :=0;i<5;i++{
		wg.Add(1)
		fmt.Printf("start goroutine %d\n",i)

		go func(i int){
			defer wg.Done()

			<-begin
			fmt.Printf("%d has begin\n",i)
		}(i)
	}

	fmt.Println("unBlocking goroutine!")
	close(begin)

	wg.Wait()

}


// チャネルは入力が終わったら閉じておかないとその後のチャネル内の値読み取りが永遠に読み込もうとしてしまいデッド路オッ苦に陥ってしまうの注意。