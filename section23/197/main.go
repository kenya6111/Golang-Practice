package main

import (
	"fmt"
	"math/rand"
)



func DoSomething() <- chan int{
	readStream := make(chan int)

	go func(){
		defer fmt.Println("DoSomething done")
		defer close(readStream)

		for{
			readStream <- rand.Intn(100)
		}
	}()

	return readStream;
}
func main() {
	readStream := DoSomething()

	for i:=1; i<=3; i++{
		fmt.Println(<-readStream)
	}
}

