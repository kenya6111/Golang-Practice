package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu sync.Mutex
	value int
}


var wg sync.WaitGroup


func main() {
	printSum := func(v1,v2 *value){
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum%v\n", v1.value + v2.value)
	}
	fmt.Println(1)

	var a,b value
	fmt.Println(2)

	wg.Add(2)
	fmt.Println(3)
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}