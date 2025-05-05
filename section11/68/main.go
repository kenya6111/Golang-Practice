package main

import "fmt"

// 独自型。MyInt型というものをintを格納できる独自の型として定義できえる
type MyInt int

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	mi += 100
	fmt.Println(mi)

	i := 100
	fmt.Println(i)
	// mi +=i MyInt とintは型が違うので演算はできない

}
