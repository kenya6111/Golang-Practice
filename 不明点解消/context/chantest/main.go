package chantest

import (
	"fmt"
	"time"
)

func TestChan() {

	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Println("chのvalue", v)
		}
	}()

	for i := 0; i < 6; i++ {
		ch <- 1
	}

	close(ch)

	ch2 := make(chan int, 3)
	fmt.Println(ch2)
	fmt.Println(len(ch2))
	fmt.Println(ch2)

	go func() {
		fmt.Println("Hello from goroutine")
	}()
	fmt.Println("Hello from main")
	time.Sleep(time.Second * 1)

}
