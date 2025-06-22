package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var lock sync.Mutex
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


	start :=time.Now()

	for i :=0; i<5; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}

	wg.Wait()
	end:= time.Now()
	fmt.Println(end.Sub(start))
}


// ✅ 実行の流れ（イメージ）
// イメージとしては、「l」が部屋みたいなイメージで、
// 1つ目の処理がその部屋をロック、
// そしてその部屋の中で1つ目の処理は処理を進捗上げることができて、
// 終わったら部屋を出る（アンロック）。
// そして次の人が入る。
// みたいなイメージ。


