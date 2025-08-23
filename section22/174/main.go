package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Hello (wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("hello from %v\n",id)

}
func main() {
	var wg sync.WaitGroup

	// wg.Add(1)
	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("1st go routine start1")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("1st go routine Done1")
	// }()

	// wg.Add(1)

	// go func () {
	// 	defer wg.Done()
	// 	fmt.Println("2st go routine start2")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("2st go routine Done2")
	// }()
	// wg.Wait()


	var CPU int = runtime.NumCPU()

	wg.Add(CPU)
	for i :=1; i<= CPU; i++{
		go Hello(&wg, i)
	}

	wg.Wait()


}
