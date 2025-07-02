package main

import (
	"fmt"
	"time"
)

func reciever (name string , ch chan int){
	for{
		i, ok := <-ch
		if !ok{
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main(){
 	ch1 := make(chan int,3)

	for i:=0; i<3;i++{
		ch1<-i
	}
	close(ch1)

	for i:= range ch1{
		fmt.Println(i)
	}
	// close(ch1)

	time.Sleep(3*time.Second)

}


