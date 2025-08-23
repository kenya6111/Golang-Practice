package main

import (
	"fmt"
	"time"
)

func generator () <-chan int{
	intStream:=make(chan int)

	go func(){
		defer close(intStream)

		for i:=0;i<=100;i++{
			intStream<-i
		}
	}()
	return intStream

}

func signal (after time.Duration) <-chan interface{}{
	done:=make(chan interface{})

	go func(){
		defer close(done)
		defer fmt.Println("signal done")
		time.Sleep(after)
	}()

	return done

}
func main() {
	done:=signal(10*time.Second)
	intStream := generator()

	loop:
	for{
		select{
			case <-done:
				break loop
			case val,ok :=<-intStream:
				if !ok{
					return 
				}
				fmt.Println(val)
		}
	}
}

// orDoneチャネル
// 終了判定のdoneチャネルと入力データのチャネルのどちらかが閉じれば閉じるチャネル

// // 
// 「or-done」という名前の通り、
// if channel1 is done OR channel2 is done OR ... then done
// という 論理和的（OR）な終了条件を構築することから来ています。


// generatorとsignalのどちらかが終了したら読み込み終了したい