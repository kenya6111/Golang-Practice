package withvalue

import (
	"context"
	"fmt"
)

// コンテキストに任意の値を入れて、後からキーで取り出す」ということを確認するテスト。
func TestWithValue() {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil { //ctx.Value(key) で値を取り出す
			fmt.Println("found value", v)
			return
		}
		fmt.Println("key not found", k)
	}
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go") //context.WithValue で値を埋め込む。つまり context に「language= GO」という情報を格納した ということ。
	f(ctx, k)
	f(ctx, favContextKey("color"))
}
