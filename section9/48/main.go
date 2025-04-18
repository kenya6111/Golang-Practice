package main

import "fmt"

func main() {
	// sl := []int{100, 200}

	// sl2 := sl
	// sl2[0] = 1000
	// fmt.Println(sl)

	// var i int = 10 // intとかの値は他の変数に代入しそれに値を入れても値は別々だが、参照型は共有されてしまう
	// i2 := i
	// i2 = 100
	// fmt.Println(i, i2)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

}
