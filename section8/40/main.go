package main

import (
	"fmt"
)

func main() {
	// ラベル付きfor
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop2
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}
