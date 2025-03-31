package main

import (
	"fmt"
	"strconv"
)

func outer() {
	fmt.Println(111)
}
func main() {
	var i int = 100
	fmt.Println(i)
	fmt.Printf("%T\n", i) // データ型がわかる。
	fmt.Printf("%T\n", int32(i))

	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) //+Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) //-Inf

	nan := zero / zero //NaN

	fmt.Println(nan)

	var t, f bool = true, false
	fmt.Println(t, f)

	fmt.Println(`test
	test
	test
	`)

	fmt.Println("\"") // ダブルクォーテーションを表示する場合。
	fmt.Println(`"`)  // ダブルクォーテーションを表示する場合。

	var sss string = "Hello golung"
	fmt.Println(sss[0])         //"H"だが72が出る
	fmt.Println(string(sss[0])) // 文字列に変換　「H」が出力

	byteA := []byte{72, 73}
	fmt.Println(byteA) // [72 73]が出る

	fmt.Println(string(byteA)) // HI が出る

	c := []byte("HI") // バイト配列に直す
	fmt.Println(c)    // [72 73]が出力される

	fmt.Println(string(c)) //

	// GOの配列型は、あとから要素数を変更できない。増減できない
	fmt.Println("----------")
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // [A B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // [1 2 3]

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)        // [C D]
	fmt.Printf("%T\n", arr4) // [2]string

	fmt.Println(arr1[0]) // 0
	fmt.Println(arr2[0]) // A

	arr2[2] = "E"
	fmt.Println(arr2[2]) // E

	fmt.Println(len(arr1)) // 3

	var x interface{} // あらゆる型と互換性がある
	fmt.Println(x)    // nil 初期値はnilとなっておりpythonでいうところのNoneになる。

	x = "aaa"
	fmt.Println(x) /// aaa
	x = 111
	fmt.Println(x) // 111
	x = true
	fmt.Println(x) // true

	x = [3]int{1, 2, 3}
	fmt.Println(x) //[1 2 3]
	// fmt.Println(x + 3)// invalid operation: x + 3 (mismatched types interface{} and int)

	var i_3 int = 1
	fmt.Printf("%T\n", i_3) //int
	fl64_2 := float64(i_3)
	fmt.Println(fl64_2)        // 1
	fmt.Printf("%T\n", fl64_2) // float64

	inti_3 := int(fl64_2)
	fmt.Println(inti_3)        // 1
	fmt.Printf("%T\n", inti_3) // int

	var s_4 string = "100"
	fmt.Printf("%T\n", s_4) // string

	i, _ = strconv.Atoi(s_4)
	fmt.Printf("%T\n", i) // int
	fmt.Print(i)          // 100

}
