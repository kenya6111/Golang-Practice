package main

import "fmt"

func init2() { //  「init」という関数は特別。mainより前に初期化処理をしたい時に使う
	fmt.Println("init")
}
func main() {
	fmt.Println("Main")
}
