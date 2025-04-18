package main

import "fmt"

func main() {
	var m = map[string]int{"a": 100, "b": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["a"])
	fmt.Println(m["aaa"])
	fmt.Println(m2["A"])
	fmt.Println(m3[1])

	s, ok := m4[1] //二つ目の変数には取り出しに成功したかどうかのやつがはいる
	fmt.Println(s, ok)

	m5 := make(map[int]string)
	fmt.Println(m4)

	m5[1] = "JAPAN"
	m5[2] = "USA"
	m5[3] = "PPP"
	m5[4] = "LAK"
	delete(m5, 3)
	fmt.Println(m5)

}
