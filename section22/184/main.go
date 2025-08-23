package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup

	// m := map[string]int{"A":0, "B":1}

	// for i:= 0; i<10; i++{
	// 	wg.Add(2)
	// 	go func (){
	// 		defer wg.Done()
	// 		m["A"] = rand.Intn(100)
	// 		m["B"] = rand.Intn(100)
	// 	}()

	// 	go func (){
	// 		defer wg.Done()
	// 		m["A"] = rand.Intn(100)
	// 		m["B"] = rand.Intn(100)
	// 	}()
	// }

	// wg.Wait()
	
	smap :=&sync.Map{}
	smap.Store("hello", "world")
	smap.Store(1,2)

	smap.Range(func(key,value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
	
	
	// smap.Delete(1)
	smap.Delete("hello")

	smap.Range(func(key,value interface{}) bool {
		fmt.Println(key,value)
		return true
	})


	smap.Store("hello", "world")

	fmt.Println("---")

	v,ok :=smap.Load("hello")

	if ok{
		fmt.Println(v)
	}

	fmt.Println("---")
	smap.LoadOrStore("hhh","wwww")// もしなければ追加する
	smap.LoadOrStore("hhh","wwww")// もしなければ追加する
	smap.LoadOrStore("hhh","wwww")// もしなければ追加する

	smap.Range(func(key,value interface{}) bool {
		fmt.Println(key,value)
		return true
	})
}

// fatal error: concurrent map writes
// concurrent：同時に、並行して（＝複数のゴルーチンが同時に）
// map writes：Goのmapに対して「書き込み操作」
// 👆これを防ぐののがsyncのmap型