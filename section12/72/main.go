package main

import "fmt"

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v,%v>>", p.A, p.B)
}
func main() {
	p := &Point{100, "ABC"} // 構造体 Point のインスタンス（実体）を作って、そのアドレス（ポインタ）を p に入れてる
	fmt.Println(p)
}
