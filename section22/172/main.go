package main

import (
	"fmt"
	"runtime"
	"sync"
)

func memConsumed () uint64 {
	runtime.GC()

	var s runtime.MemStats

	runtime.ReadMemStats(&s)

	return s.Sys
}

func main() {
	var ch <-chan interface{}
	var wg sync.WaitGroup

	noop := func(){
		wg.Done()
		<-ch
	}
	const numGoroutines = 1000000

	wg.Add(numGoroutines)

	before := memConsumed()

	for i := 0; i< numGoroutines; i++{
		go noop()
	}

	wg.Wait()

	after:= memConsumed()

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

}
