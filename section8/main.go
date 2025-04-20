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

func anything(a interface{}) {
	fmt.Println(a)

}
func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 { //if や else ifには 丸括弧いらない
		fmt.Println("i dont know")
	} else {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true { // この行のxはスコープはif内なので外のxには干渉しない
		fmt.Println(x)
	}

	fmt.Println(x)

	// エラーハンドリング
	var s string = "100"
	i_8, _ := strconv.Atoi(s)
	fmt.Println(i_8)
	fmt.Printf("%T", i_8)
	fmt.Println()

	var s2_2 string = "das"
	i_8_2, err := strconv.Atoi(s2_2) // errには変換できなかった時のエラー分が返ってくる
	fmt.Println(err)
	fmt.Println(i_8_2)
	fmt.Printf("%T", i_8_2)
	fmt.Println()

	// for文 へえ〜ってなったものだけ
	// 条件なしfor
	i := 0
	for {
		i++
		if i == 10 {
			break // for文を強制終了させるやつ
		}
		fmt.Println(i)
	}

	fmt.Println("----")

	// 条件付きfor
	i_2 := 0
	for i_2 < 10 {
		i_2++
		fmt.Println(i_2)
	}

	fmt.Println("----")
	// 古典的for
	for i_3 := 0; i_3 < 10; i_3++ {
		fmt.Println(i_3)

	}
	fmt.Println("----")

	arr2 := []int{111, 222, 333, 444, 555, 666, 777}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	fmt.Println("----")

	// 範囲式
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("----")

	sl := []string{"python", "golang", "java"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	fmt.Println("----")
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// switch
	n := 11
	switch n {
	case 1, 2:
		fmt.Println(111)
	case 3, 4:
		fmt.Println(222)
	case 11:
		fmt.Println(333)
	default:
		fmt.Println(444)

	}

	switch n := 2; n { // switch文ないのみで参照可能な書き方
	case 1:
		fmt.Println(111)
	case 2:
		fmt.Println(222)
	default:
		fmt.Println(333)

	}

	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("dasdasdsad")
	case n2 > 3 && n2 < 7:
		fmt.Println("dasdasdsaddasadasdas")
	}

	// switch　型スイッチ
	anything(1)
	anything(1)
	anything(true)

	// アサーションはただ型を判定してるだけのやつ。キャストみたいに変換はしていない
	var x3 interface{} = 3
	i3 := x3.(int) // 型のアサーション 変数.(復元したい型)でキャスト的なことができる
	fmt.Println(i3 + 2)

	f2, isFloat64 := x3.(float64) // 2つ目の変数には変換成功したかどうかのしんぎちが入る
	fmt.Println(f2, isFloat64)

	var x4 interface{} = "3"
	if x4 == nil {
		fmt.Println("None")
	} else if i, isInt := x4.(int); isInt {
		fmt.Println("x is int")
		fmt.Println(i)
	} else if s, isString := x4.(string); isString {
		fmt.Println("x is string")
		fmt.Println(s)
	} else {
		fmt.Println("i dont know")
	}

	var x5 interface{} = 5
	i5, isInt := x5.(int)
	fmt.Println(i5, isInt)
}
