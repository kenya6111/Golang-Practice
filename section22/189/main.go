package main

import (
	"fmt"
	"time"
)

func main() {
	ch :=make(chan int)
	go func(){
		defer fmt.Println("chan closed")
		defer close(ch)
		for i:=0; i<5; i++{
			fmt.Printf("writing to channel: %v\n", i)
			ch <-i
		}
	}()

	for integer := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("reading to chan %v\n",integer)
	}
}