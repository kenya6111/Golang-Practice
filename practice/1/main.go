package main

import (
	"fmt"
	"time"
)


func main(){
	start := time.Now()
	user := fetchUser()
	followers := fetchFollowers(user)
	posts := fetchPosts(user)
	fmt.Println("followers:", followers)
	fmt.Println("posts:", posts)
	fmt.Println("time: ", time.Since(start))
}


func fetchUser() string{
	time.Sleep(time.Millisecond*50)
	return "tanaka tarou"
}
func fetchFollowers(user string) int{
	time.Sleep(time.Millisecond*50)
	return 1234
}

func fetchPosts(user string)[]string{
	time.Sleep(time.Millisecond * 100)
	return []string{"GoでWeb API作った", "Docker勉強中", "MySQL接続で詰まった"}
}