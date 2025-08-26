package main

import "fmt"

func main() {
	// &で任意の型からそのポインタ型を生成できる
	var i int
	i = 10                 // 10
	fmt.Printf("%T\n", i)  // int
	fmt.Printf("%T\n", &i) // *int

	ii := &i
	iii := &ii
	fmt.Printf("%T\n", ii)  // *int
	fmt.Printf("%T\n", iii) // **int

	fmt.Println("--------")
	var a int
	p := &a
	a = 5
	fmt.Println(*p)
	*p = 10
	fmt.Println(a)
	fmt.Println("--------")

	b := 1
	inc(&b)
	inc(&b)
	inc(&b)
	fmt.Println(b) // 4

	fmt.Println("--------")
	c := &[3]int{1, 2, 3}
	inc2(c)
	inc2(c)
	inc2(c)
	fmt.Println(c)
	fmt.Printf("%p\n", c)

	fmt.Println("--------")

}

// *[データ型]でポインタ型を受け取れる型にできて
// 変数に*を使うと、そのポインタの値を取れるってこと
func inc(p *int) {
	*p++
}
func inc2(p *[3]int) {
	for i := 0; i < 3; i++ {
		(*p)[i]++
	}
}
