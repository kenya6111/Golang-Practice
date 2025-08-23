package main

import withvalue "resolve/context/withValue"

func main() {
	// afterfunc.TestAfterFunc()
	// cause.TestCause()
	// n := 0 // ちょっとfouぶん練習
	// for {
	// 	if n == 100 {
	// 		break
	// 	}
	// 	n++
	// }
	// fmt.Println(n)
	// withcancel.TestWithCancel()

	// donetest.TestDone2()
	// canceltest.TestCancel4()

	// deadline.TestDeadline2()
	withvalue.TestWithValue()
	// chantest.TestChan()
	// fmt.Println(212)
	// fmt.Println(context.Background())
	// fmt.Printf("aaa: %T\n", context.Background())
	// fmt.Printf("aaa: %T\n", 333)

	// // withValue
	// type favContextKey string

	// f := func(ctx context.Context, k favContextKey) {
	// 	if v := ctx.Value(k); v != nil {
	// 		fmt.Println("found value:", v)
	// 		return
	// 	}
	// 	fmt.Println("key not found:", k)
	// }

	// k := favContextKey("language")
	// ctx := context.WithValue(context.Background(), k, "Go")

	// f(ctx, k)
	// f(ctx, favContextKey("color"))

	// c := context.Background()
	// ctx2 := context.WithValue(c, "userID", 42)
	// ctx2 = context.WithValue(ctx2, "testKey", false)
	// fmt.Println(ctx2)
	// fmt.Println(ctx2.Value("testKey"))
	// fmt.Println(ctx2.Value("userID"))

	// // withCancel
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx = context.WithValue(ctx, "userID", 999)

	// go func() {
	// 	<-ctx.Done()
	// 	fmt.Println("canceled:", ctx.Err())
	// }()

	// cancel()

	// fmt.Println(ctx.Value("userID"))

	// // context.Done()
	// ctx4, cancel4 := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel4()

	// go func() {
	// 	fmt.Println("Waiting for cancel...")
	// 	<-ctx4.Done() // ここでキャンセルを待つ
	// 	fmt.Println("Context canceled:", ctx4.Err())
	// }()

	// time.Sleep(3 * time.Second)
	// fmt.Println("main done")

}
