package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	fmt.Println(math.Abs(-33))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Cbrt(8))
	fmt.Println(math.Max(11, 22))

	fmt.Println(math.Trunc(math.Pi))
	fmt.Println(math.Trunc(3.456))
	fmt.Println(math.Trunc(3.956))
	fmt.Println(math.Trunc(4.026))
	fmt.Println(math.Floor(1.51))
	fmt.Println(math.Floor(1.31))
	fmt.Println(math.Floor(2.31))
	fmt.Println(math.Floor(2.01))
	fmt.Println(math.Floor(1.91))
	fmt.Println(math.Ceil(1.91))

}
