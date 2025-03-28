package main

import "fmt"

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

	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr1[0])
	fmt.Println(arr2[0])

	arr2[2] = "E"
	fmt.Println(arr2[2])

	fmt.Println(len(arr1))

	// 	【Goの配列型】
	// 他のプログラミング言語の配列型になれていると、若干奇妙かもしれませんが、Goの配列型は要素数を変更できません。

	// 要素の追加などを行う場合は、後のレクチャーで登場するスライス型を使います。

	// 配列型　＝　要素数を変更できない。

	// スライス型　＝　要素数を変更可能。

	var x interface{} //インタフェース方は全ての型と互換性がある。xにあらゆる型のデータを入れれる
	fmt.Println(x)    // <nil>が返される。PythonでいうところのNoneと同じ
	// だが普通の演算はできない型になる
	x = 1
	fmt.Println(x)

	x = "aaa"
	fmt.Println(x)

	// 型の変換
	var i_2 int = 1
	fl64_2 := float64(i_2)
	fmt.Println(fl64_2)
	fmt.Printf("i = %T\n", i_2)
	fmt.Printf("fl64 = %T\n", fl64_2)

	i5 := int(fl64_2)
	fmt.Print("i5 = %T\n", i5)

	fl6 := 10.5
	i6 := int(fl6)

	fmt.Print("i6= %T\n", i6)
	fmt.Print(i6, "\n")

	// 試行錯誤らん
	fmt.Printf("%Tと%Tと%T", 1, 2, 3)
	fmt.Print("\n")
	fmt.Print(1, 2, 3)
	fmt.Print("\n")

	name := "ALice"
	age := 30
	fmt.Printf("名前は%vです。年齢は%vです", name, age)

	fmt.Print("\n")
	var sss2 string = "abcABCHello"
	fmt.Println(sss2[0])
	fmt.Println(sss2[1])
	fmt.Println(sss2[2])
	fmt.Println(string(sss2[0]))
	fmt.Println(string(sss2[1]))
	fmt.Println(string(sss2[2]))
	fmt.Println(sss2[3])
	fmt.Println(sss2[4])
	fmt.Println(sss2[5])
}
