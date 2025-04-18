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

	var p *int = &n // アスタリスクでポインタ型を宣言。「&nでアドレスを渡している」

	fmt.Println(p)
	fmt.Println(*p) // ポインタ型の値を見るには＊を変数の前につける

	*p = 300
	fmt.Println(n)

	double2(&n) // &をつけることで参照を渡している。
	fmt.Println(n)

	double2(p)
	fmt.Println(*p)

}
