package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v,name=%v", p.Name, p.Age)
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v,Model=%v", c.Number, c.Model)
}

type Stringfy interface {
	ToString() string
}

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
