package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
}
func main() {

	// myPool := &sync.Pool{
	// 	New : func () interface{}{
	// 		fmt.Println(" create  new person instance")
	// 		return new(Person)
	// 	},
	// }

	// myPool.Put(&Person{Name:"1"})
	// myPool.Put(&Person{Name:"2"})

	// instance1 := myPool.Get()
	// instance2 := myPool.Get()
	// instance3 := myPool.Get().(*Person)

	// fmt.Println(instance1,instance2,instance3)
	// instance3.Name="3"
	// fmt.Println(instance1,instance2,instance3)

	// myPool.Put(instance1)
	// myPool.Put(instance2)
	// myPool.Put(instance3)

	// instance4 := myPool.Get()
	// instance5 := myPool.Get()
	// instance6 := myPool.Get()
	// fmt.Println(instance4,instance5,instance6)
	// fmt.Println(myPool.Get())


	// ppolでどれくらいのインスタンスの生成を節約できるか　
	count := 0
	myPool := &sync.Pool{
		New : func() interface{}{
			count++
			fmt.Println("Creating...")
			return struct{}{}
		}, 
	}

	myPool.Put("manually added: 1")
	myPool.Put("manually added: 2 ")

	var wg sync.WaitGroup

	wg.Add(10000)

	for i:=0; i<10000; i++{
		time.Sleep(1*time.Millisecond)
		go func(){
			defer wg.Done()
			instance := myPool.Get()
			myPool.Put(instance)
		}()
	}
	wg.Wait()
	fmt.Printf("created instance %d", count)
}

// pool
// オブジェクトのキャッシュを作るやつ
// Put()でオブジェクトをプールにおける
// Get()でputしたオブジェクトぶんだけプールから取得できる。、取れる順番はキューの構造となっている
