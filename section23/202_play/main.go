package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func random() interface{}{
	return rand.Intn(5000000000)

}
func repeatFunc (done <-chan interface{}, fn func() interface{}) <-chan interface{}{
	valueStream :=make(chan interface{})
	go func(){
		defer close(valueStream)
		for{
			select{
			case <- done:
				fmt.Println("doneがクローズ")
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int{
	intStream:= make(chan int)

	go func(){
		defer close(intStream)

		for v:=range valueStream {
			select{
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()

	return intStream
}

func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{}{
	primeStream:= make(chan interface{})
	go func(){
		defer close(primeStream)
		L:
		for i:=range intStream {
			for div := 2; div<i; div++{
				if i%div==0{
					continue L
				}
			}
			select {
			case <-done:
				return
			case primeStream <-i:
			}
		}
	}()
	return primeStream
}
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{}{
	takeStream:= make(chan interface{})
	go func(){
		defer close(takeStream)
		for i:=0; i<num; i++ {
			select{
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func main() {
	done := make(chan interface{})

	defer close(done)

// fmt.Println(rand.Intn(5000000000))
// fmt.Println(rand.Intn(5000000000))
// fmt.Println(rand.Intn(5000000000))

	fmt.Println(repeatFunc(done,random))
	// valueStream := repeatFunc(done,random)
	randIntStream := toInt(done,repeatFunc(done,random))

	count:=0
	for v := range randIntStream{
		fmt.Println(v)
		count++

		if count >100{
			break
		}
	}

	numFinders := runtime.NumCPU()
	fmt.Printf("prime finders: %v\n", numFinders)

	finders := make([]<-chan interface{},numFinders)
	fmt.Println(finders)
	
	for i:=0;i<numFinders;i++{
		finders[i] = primeFinder(done, randIntStream)
	}
	fmt.Println(finders)
}


// return は**「forを抜ける」のではなく、「関数そのものを終了する」**
// break	ループだけを抜ける	一番内側の for / switch / select