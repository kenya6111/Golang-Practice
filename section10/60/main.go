package main

import "fmt"

func double(i int) {
	i = i * 2
}

func double2(i *int) {
	*i = *i * 2
}
func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) // メモリのアドレスを取得

	double(n)
	fmt.Println(n) // intなどは値型と呼ばれていて、関数の引数と渡してもコピーされて渡されるので元の引数とは別物になる

	var p *int = &n // アスタリスクでポインタ型を宣言。「&nでアドレスを渡している」 ＆をつけることでアドレスに変わるっぽい？

	fmt.Println(p)  // &nでアドレス値を渡しているので、実態の値を見るには⇩
	fmt.Println(*p) // ポインタ型の値を見るには＊を変数の前につける

	*p = 300
	fmt.Println(n)

	double2(&n) // &をつけることで参照を渡している。
	fmt.Println(n)

	double2(p)
	fmt.Println(*p)

	var n2 int = 10
	fmt.Println(n2)
	fmt.Println(&n2)

	fmt.Println("----")
	var n3_adr int = 11
	var n3 *int = &n3_adr // アドレス渡す
	fmt.Println(n3)
	*n3 = 11
	fmt.Println(n3)

	fmt.Println("----")
	x := 10  // 普通のint型変数
	p3 := &x // pはxのアドレスを持ってるポインタ

	fmt.Println(p3)  // アドレス表示 → 0xc000014090 とか
	fmt.Println(*p3) // pが指してるxの中身 → 10

	// 実験
	nn := 200
	fmt.Println(nn)
	fmt.Println(&nn) // アドレス取得

	var mm *int = &nn        //アドレスをmmに代入
	fmt.Println("mm:", mm)   //mmにはアドレスが入っている
	fmt.Println("*mm:", *mm) // アドレスの変数は「＊」で値を取得
	fmt.Println(Test(nn))    //
	fmt.Println(nn)
	fmt.Println(Test2(mm))
	fmt.Println(nn)

	sl := []int{1, 2, 3, 4, 5}
	Test3(sl)
	fmt.Println(sl)

}

func Test(num int) int {
	num = num * 2
	return 2 * num
}

// アドレスを渡すことで元の値も更新する
func Test2(num *int) int {
	*num = *num * 2
	return *num
}

// スライスなどの参照型はデフォルトで参照私となるため「＊」や「＆」などいちいち書かなくていい
func Test3(sl []int) {
	for i, v := range sl {
		sl[i] = v * 2
	}
}

// * は「ポインタ（＝アドレスを持ってる変数）」に使って、
// そのアドレスが指してる中身（実体の値）にアクセスするための演算子！

// 使われ方           | 意味                 | 例
// 型の定義で使う      | 「この変数はポインタ型だよ」 | var p *int（int型へのポインタ）
// 値を取り出すときに使う | 「ポインタの中身を取り出せ」 | *p = 10、fmt.Println(*p)

// 書き方 | 読み方
// var p *int | 「int型のアドレスを入れる箱を作る」
// *p = 100 | 「そのアドレスが指す先に100入れる」
// fmt.Println(*p) | 「アドレスの先の中身を表示する」

// var p *int → 型宣言で「これはポインタ」
// *p → 実体参照で「ポインタの中身を触る」

// & →「変数のアドレス（住所）を取得する演算子
// x := 42
// p := &x
