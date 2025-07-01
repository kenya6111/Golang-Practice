package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	start := time.Now()
	ch := make(chan any,2)

	var wg sync.WaitGroup
	wg.Add(2)

	user := fetchUser()
	go fetchFollowers(user,ch, &wg)
	go fetchPosts(user,ch,&wg)

	wg.Wait()

	close(ch)
	for res := range ch{
		fmt.Println(res)
	}
	fmt.Println("time: ", time.Since(start))
}


func fetchUser() string{
	time.Sleep(time.Millisecond*50)
	return "tanaka tarou"
}
func fetchFollowers(user string,ch chan any, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Millisecond*50)
	fmt.Println("fetch followers done")
	ch <-1234
}

func fetchPosts(user string, ch chan any, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println("fetch posts done")
	ch<-[]string{"GoでWeb API作った", "Docker勉強中", "MySQL接続で詰まった"}
}