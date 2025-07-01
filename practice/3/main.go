package main

import "fmt"


func main(){

	// 通常のforループ
	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }

	// 条件だけのforループ
	// i:=0
	// for i<5{
	// 	fmt.Println(i)
	// 	i++
	// }

	// 無限ループ
	// for {
	// 	fmt.Println(123)
	// }


	// range を使ったループ
	// nums := []int{1,2,3,4,5,6}
	// for i,v := range nums {
	// 	fmt.Println(i,v)
	// }
	// for a:= range nums {
	// 	fmt.Println(a)
	// }
	// for _,b:= range nums {
	// 	fmt.Println(" ",b)
	// }


	// breakを使ったやつ
	// for i:=0; i < 5;i++{
	// 	if i==3{
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }


	// 2重forでのbreak →今のforを全部飛ばすって役割
	// for i:=0; i < 5;i++{
	// 	fmt.Println("--i--: ",i)
	// 	for j:=0; j<5;j++{
	// 		if j==3{
	// 			break
	// 		}
	// 		fmt.Println("J: ",j)
	// 	}
	// }

	// 今のforの今の周回だけとばす。
	// for i:=0; i < 5;i++{
	// 	fmt.Println("--i--: ",i)
	// 	for j:=0; j<5;j++{
	// 		if j==3{
	// 			continue
	// 		}
	// 		fmt.Println("J: ",j)
	// 	}
	// }

	// break ラベル名でそこまで飛ばす
	// Loop:
	// for i:=0; i < 5;i++{
	// 	fmt.Println("--i--: ",i)
	// 	for j:=0; j<5;j++{
	// 		if j==3{
	// 			break Loop
	// 		}
	// 		fmt.Println("J: ",j)
	// 	}
	// }

	outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
	// i: 0 j: 0
	// i: 1 j: 0
	// i: 2 j: 0

	// outer:
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		if j == 1 {
	// 			break outer
	// 		}
	// 		fmt.Println("i:", i, "j:", j)
	// 	}
	// }
	//i: 0 j: 0


	// break outer	ラベル付き for を完全に抜ける
	// break outer	すぐ outer: の外に出る	ループ強制終了


	// continue outer	ラベル付き for の次のイテレーションへ進む
	// continue outer	outer: の次の回へすぐ飛ぶ	ループスキップして継続

	// break outer：
	// 	条件が成立した時点で、全部やめたいとき
	// 	例：検索して見つかったらそれ以上見なくていい

	// continue outer：
	// 	内側の条件次第で、今の外ループは無効にして次に進みたいとき
	// 	例：フィルタやバリデーション





}