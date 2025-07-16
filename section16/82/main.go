package main

import (
	"flag"
	"fmt"
)

func main() {
	// ① フラグを宣言（戻り値は **ポインタ**）
	name := flag.String("name", "guest", "ユーザー名")
	age := flag.Int("age", 0, "年齢")
	debug := flag.Bool("debug", false, "デバッグモード")

	// ② 解析：これ以降 os.Args が処理済みに
	flag.Parse()

	// ③ 使う
	fmt.Printf("name=%s age=%d debug=%v\n", *name, *age, *debug)

}
