package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// https://pkg.go.dev/sync

var once sync.Once
var once2 sync.Once

func onceMethod() {
	fmt.Println("この処理は一回のみ実行")
}
func main() {
	// sync.OnceのメソッドDo
	for i := 0; i < 3; i++ {
		once.Do(onceMethod)
	}
	for i := 0; i < 3; i++ {
		once2.Do(func() {
			fmt.Println("この処理は1回だけ実行される")
		})
	}

	// syncの関数のOnceFunc
	oncefun := sync.OnceFunc(func() { fmt.Println("oncefunc!!") })
	oncefun()
	oncefun()
	oncefun()

	// syncの関数のonceValue. 最初の実行結果だけを保存して配る装置
	onceval := sync.OnceValue(func() int {
		var num = 1000
		return num
	})

	fmt.Println(onceval())
	fmt.Println(onceval())

	initVal := sync.OnceValue(func() string {
		fmt.Println("初期化中...")
		fmt.Println("初期化中...")
		return "expensive result"
	})

	fmt.Println(initVal())
	fmt.Println(initVal())
	fmt.Println(initVal())

	once := sync.OnceValue(func() int {
		sum := 0
		for i := 0; i < 1001; i++ {
			sum += i
		}
		fmt.Println("Computed once:", sum)
		return sum
	})
	// once()
	// once()
	// once()

	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			const want = 499500
			got := once()
			if got != want {
				fmt.Println("want", want, "got", got)
			}
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}

	// syncの関数のonceValues
	once2 := sync.OnceValues(func() ([]byte, error) {
		fmt.Println("Reading file once")
		return os.ReadFile("example_test.go")
	})
	done2 := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			data, err := once2()
			if err != nil {
				fmt.Println("error:", err)
			}
			_ = data // Ignore the data for this example
			done2 <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done2
	}

	// syncのtypeのcond
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("goroutine %d: wait開始\n", id)
			cond.Wait() // ← ここで待機状態になる
			fmt.Printf("goroutine %d: wait解除\n", id)
			mu.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main: Broadcastで全員を起こす")
	cond.Broadcast()
	wg.Wait()
	fmt.Println("完了")

	// condのSignal
	var mu3 sync.Mutex
	cond3 := sync.NewCond(&mu3)
	var wg3 sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// 3つの goroutine が待機
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			mu3.Lock()
			fmt.Println("goroutine", id, "wait開始")
			cond3.Wait()
			fmt.Println("goroutine", id, "wait解除")
			mu3.Unlock()
		}(i)
	}
	// 少し待ってから1つずつ起こす
	time.Sleep(time.Second)
	for j := 0; j < 3; j++ {
		mu3.Lock()
		fmt.Println("main: Signalで1人だけ起こす")
		cond3.Signal()
		mu3.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	wg3.Wait()
}
