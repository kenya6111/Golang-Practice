package main

import (
	"fmt"
	"time"
)

// func or(channels ...<-chan interface{}){

// }


func generateChan(after time.Duration) chan interface{}{
	done := make(chan interface{})

	go func () {
		defer close(done)
		time.Sleep(after)
	}()

	return done
}

func main() {
	// start := time.Now()
	// time.Sleep(2*time.Second)
	// fmt.Printf("done after: %v\n", time.Since(start))
	// chan1 := generateChan(4 * time.Second)
	// chan2 := generateChan(4 * time.Minute)
	// chan3 := generateChan(4 * time.Hour)
	// // <-chan1 // 🔥 ここで4秒待つ
	// fmt.Println(chan1)
	// fmt.Println(chan2)
	// fmt.Println(chan3)


	// i := 3

    // switch i {
    // case 0:
    //     fmt.Println("i is 0")
    // case 1:
    //     fmt.Println("i is 1")
    // default:
    //     fmt.Println("x is neither 1 nor 2")
    // }

	// switch {
    // case i == 0:
    //     fmt.Println("i is 0")
    // case i == 1:
    //     fmt.Println("i is 1")
    // default:
    //     fmt.Println("x is neither 1 nor 2")
    // }

	// x := 4
	// switch x {
    // case 1, 2, 3:
    //     fmt.Println("x is either 1, 2, or 3")
    // case 4, 5, 6:
    //     fmt.Println("x is either 4, 5, or 6")
    // default:
    //     fmt.Println("x is not 1, 2, 3, 4, 5, or 6")
    // }

	// 容量 0
	// ch1 := make(chan int)
	// // 容量 10
	// ch2 := make(chan int,10)

	// ch := make(chan int)
	// go func (){
	// 	fmt.Println(<-ch)
	// }()
	// go func (){
	// 	fmt.Println(<-ch)
	// }()
	// ch<-1
	// ch<-2
	// close(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	
	s := []int{1, 2, 3}
    s = append(s, 4, 5, 6)
    // fmt.Println(s)
    // fmt.Println(s[0])
    // fmt.Println(s[2])
    // fmt.Println(s[:1])
    // fmt.Println(s[0:4])
    // fmt.Println(s[1:4])
    // fmt.Println(s[:4])
    // fmt.Println(s[3:])
    fmt.Println(append(s[3:],99,99,99))
    fmt.Println(append(s[2:],99,99,99))
    fmt.Println(append(s[1:],99,99,99))
    fmt.Println(append(s[6:],99,99,99))
	
	// s := []string{"サーモン", "たまご", "まぐろ"}
	// eat(s...)
}

// func eat(sushi ...string) {
// 	fmt.Println(sushi)
// }


// eat(s)    // ❌ 型エラー：スライスを渡してる



// switch
// channelsのlenとcap
// closeしたらどうなるのか
// append [3:]


// ...