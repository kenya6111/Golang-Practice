package main

import (
	"fmt"
	"reflect"
)

func main() {
	// sl := []interface{}{1, "2", 10, "11"}// interface型は、「interface{}」と書くので、先はNG。
	sl := []interface{}{1, "2", 10, "11"}
	for _, v := range sl {
		// fmt.Printf("%T\n", v)
		fmt.Println(reflect.TypeOf(v))

		fmt.Printf("%T\n", reflect.TypeOf(v))
		// fmt.Println(v)

		// fmt.Println(fmt.Sprintf("%d", v))
	}
}
