package main

import "fmt"

func generator (done <-chan interface{}, integer ...int) <-chan int{
	intStream := make(chan int)

	go func(){
		defer close(intStream)

		for _,v := range integer{
			select{
			case <-done:
				return
			case intStream <-v:
			}
		}
	}()

	return intStream
}

func double (sl []int) []int{
	doubleSlice := make([]int, len(sl))

	for i, v := range sl{
		doubleSlice[i] = v * 2
	}

	return doubleSlice
}


func add(sl []int) []int{
	addSlice := make([]int , len(sl))

	for i, v := range sl{
		addSlice[i] = v + 1
	}

	return addSlice
}
func main() {

	ints := []int{1,2,3,4,5}

	for _, v :=range double(add(ints)) {
		fmt.Println(v)

	}
}


// パイプライン