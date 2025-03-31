package main

import "fmt"

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
	i := Plus(1, 2)
	fmt.Println(i)
	i2 := Plus2(1, 2)
	fmt.Println(i2)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i5 := Double(9)
	fmt.Println(i5)
	NOReturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i4 := f(9, 9)
	fmt.Println(i4)

	i6 := func(x, y int) int {
		return x + y
	}(4, 5)
	fmt.Println(i6)

	// 関数を返す関数
	f2 := Returnfunc()
	f2()

	// 関数を引数にとる関数
	CallFUnction(func() {
		fmt.Println("i am function")
	})

	// クロージャー
	f3 := Later()
	fmt.Println(f3("hello world"))
	fmt.Println(f3("hello world2"))
	fmt.Println(f3("hello world3"))
	fmt.Println(f3("hello world4"))
	fmt.Println(f3("hello world5"))
	fmt.Println(f3("hello world6"))

	// ジェネレーター(何らかのルールに従って連続した値を返し続ける)
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

}
