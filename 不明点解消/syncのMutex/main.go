package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mu.Lock() // 🔒 ロック開始
			counter++ // 共有変数に安全にアクセス
			fmt.Println(id, counter)
			mu.Unlock() // 🔓 ロック解放
		}(i)
	}

	wg.Wait()
	fmt.Println("最終:", counter)
}
