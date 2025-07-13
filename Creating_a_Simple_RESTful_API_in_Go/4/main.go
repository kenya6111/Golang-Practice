package main

import "fmt"

func main(){
	fmt.Println("go test")
	result := Calculate(2)
	fmt.Println(result)
}

func Calculate (x int)(result int){
	result = x+2
	return result
}