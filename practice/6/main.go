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
 	ch1 := make(chan int)


	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	for i:=0; i<100;i++{
		ch1<-i
	}
	close(ch1)

	time.Sleep(3*time.Second)

}


