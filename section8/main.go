package main

import (
	"fmt"
	"strconv"
)

func Plus(x int, y int) int {
	return x + y

}
func Plus2(x, y int) int {
	return x + y

}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) { // return resultと書かなくてもここに返す変数を書けば返せる
	result = price * 2
	return
}

func NOReturn() {
	fmt.Println("no return ")
}

func Returnfunc() func() {
	return func() { fmt.Println("asaaaaa") }
}

func CallFUnction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("i dont know")
	} else {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}

	fmt.Println(x)

	// エラーハンドリング
	var s string = "100"
	i_8, _ := strconv.Atoi(s)
	fmt.Printf("%T", i_8)
	fmt.Println()

	// for文 へえ〜ってなったものだけ
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	sl := []string{"python", "golang", "java"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// switch
}
