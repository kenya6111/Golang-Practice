package main

import "fmt"

type Person struct { // OK
	Name string
	Age  int
}

type Car struct { // OK
	Number string
	Model  string
}

// ：なぜ「レシーバ（receiver）」と呼ばれるのか？
// それは、その関数を「受け取る（= receive）」型を指定するから

// *Person 型のポインタレシーバを持つ ToString メソッド と呼べる
func (p *Person) ToString() string { // OK
	return fmt.Sprintf("Name=%v,name=%v", p.Name, p.Age)
}

func (c *Car) ToString() string { // OK
	return fmt.Sprintf("Number=%v,Model=%v", c.Number, c.Model)
}

type Stringfy interface {
	ToString() string
}

// interface とは？
// → 「こういうメソッドを持ってる型だけ通してOK！」という「型のルール」

// type MyInt int      // int をベースにした新しい型 MyInt を作る
// type User struct{}  // struct の定義に User という名前をつける
func main() {

	vs := []Stringfy{
		&Person{Name: "Taro", Age: 20},
		&Car{Number: "T111-22", Model: "modemodemdoe"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}

// people := []*Person{
// 	{Name: "Taro", Age: 20},
// }
// cars := []*Car{
// 	{Number: "T111-22", Model: "modemodemdoe"},
// }

// for _, p := range people {
// 	fmt.Println(p.ToString())
// }

// for _, c := range cars {
// 	fmt.Println(c.ToString())
// }
// 配列が []Stringfy じゃなくて、[]*Person と []*Car に分かれてしまう

// 「Person と Car をまとめて扱う」ことができない
