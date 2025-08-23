package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {

	fmt.Println("ssss")
	fmt.Println(http.MethodGet)
	fmt.Println(http.MethodPost)
	fmt.Println(http.StatusAlreadyReported)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello world") // io.WriteStringになぜw,と"helloworld"を渡すと、画面位それが返ってくるのか
	}

	http.HandleFunc("/hello", helloHandler)

	if _, err := strconv.Atoi("a"); err != nil {
		// e := err.(*strconv.NumError)
		// fmt.Println("Func:", e.Func)
		// fmt.Println("Num:", e.Num)
		// fmt.Println("Err:", e.Err)
		// fmt.Println("Err::::", e.Err.Error())
		// fmt.Println(err)
		// fmt.Printf("%T\n", err)
		// fmt.Printf("err0: [%T] %v\n", err, err)
		err1 := errors.Unwrap(err)
		fmt.Println("---")
		fmt.Println(err1)
		fmt.Println("")
		fmt.Printf("[%T],%v", err1, err1)
		fmt.Println("---")
		err2 := errors.Unwrap(err1)
		fmt.Printf("err2: [%T] %v\n", err2, err2)
	}

	client := new(http.Client)

	_, err := client.Get("ggggggg")
	fmt.Println("err.Error()", err.Error())
	e := err.(*url.Error)
	fmt.Println("Op:", e.Op)   // Op: Get
	fmt.Println("URL:", e.URL) // URL: fooooo
	fmt.Println("Err:", e.Err) // Err: unsupported protocol scheme ""
	fmt.Println(err)
	log.Fatal(http.ListenAndServe(":8089", nil)) // log.Fatalって何

}
