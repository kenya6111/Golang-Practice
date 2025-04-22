package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// sl := []interface{}{1, "2", 10, "11"}// interface型は、「interface{}」と書くので、先はNG。
	sl := []interface{}{1, "2", 10, "11"}
	for _, v := range sl {

		// fmt.Printf("%T\n", v)
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.String {
			// fmt.Printf("%02d\n", v)

			i, _ := strconv.Atoi(v.(string))
			fmt.Printf("%02d\n", i)
		}
		if rv.Kind() == reflect.Int {
			fmt.Println(v)
		}
	}
}
