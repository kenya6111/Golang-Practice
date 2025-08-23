package main

import "fmt"

type Person struct {
	name string
}

type Animal struct {
	name string
}

func (p Person) Greet() string {
	return "hello ,myname is " + p.name
}
func (p *Person) changeName() {
	p.name = "change!"
}

func (p Person) changeName2() {
	p.name = "change!!!!!!!!!!!!!"
}

func changeName(a *Animal) {
	a.name = "change name!!!"
}
func main() {
	fmt.Println(11)

	p := Person{name: "kenya"}

	fmt.Println(p.name)
	fmt.Println(p.Greet())
	p.changeName()
	fmt.Println(p.name)
	p.changeName2()
	fmt.Println(p.name)

	dog := Animal{
		name: "komugi",
	}

	changeName(&dog)
	fmt.Println(dog)

}

// &（アンパサンド）はその変数からポインタを抽出するのに用いる
// * はポインタ型を示す → なので、メソッドの引数に*が付与されている場合は、ポインタ型の変数を引数として受け取るということを示しています。
