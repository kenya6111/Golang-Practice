package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var lock sync.RWMutex
	var wg sync.WaitGroup

	increment := func(wg *sync.WaitGroup, l sync.Locker){
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("increment!")
		count++
		fmt.Println("increment結果: ", count)
		time.Sleep(1 * time.Second)
	}

	read := func(wg *sync.WaitGroup, l sync.Locker){
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("read: ",count)
		time.Sleep(1 * time.Second)
	}

	start :=time.Now()

	fmt.Println("------こっからincrement-------")
	for i :=0; i<5; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}

	fmt.Println("------こっからread-------")
	for i :=0; i < 5; i++ {
		wg.Add(1)
		go read(&wg, lock.RLocker())
	}

	wg.Wait()

	end:= time.Now()

	fmt.Println(end.Sub(start))
}

// Mutex →読み込みと書き込みの両方をロックできる
// RWMutex →読み込みだけであればロックの解放を待たずに処理できる。書き込みはロックされる