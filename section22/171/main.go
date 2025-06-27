package main

import (
	"fmt"
	"sync"
)

// func sayHello (wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println("hello world")
// }
func main() {
	var wg sync.WaitGroup
	// tasks := []string{"A","B","C"}
	// for _,task := range tasks{
	// 	wg.Add(1)

	// 	go func(){
	// 		defer wg.Done()
	// 		fmt.Println(task)
	// 	}()
	// }


	// wg.Wait()

	// fruits := []string{"🍎 Apple", "🍌 Banana", "🍇 Grape", "🍊 Orange"}

	// for _,v := range fruits{
	// 	wg.Add(1)

	// 	go func (){
	// 		defer wg.Done()
	// 		fmt.Println(v)
	// 	}()

	// }

	// wg.Wait()
	// fmt.Println("✅ All tasks done!")



	say := "hello"

	wg.Add(1)

	go func (){
		defer wg.Done()
		say = "Good bye"
	}()

	wg.Wait()

	fmt.Println(say)



}


// go func(){}を実行するとメイン関数が先に終わって何も出力されない → なので終了動機が必要
// go メソッドの直前でAdd()することが推奨されている

