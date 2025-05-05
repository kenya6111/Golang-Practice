package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("start")
	// fmt.Println(math.Pow(2, 2))

	// _, err := os.Open("aa.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length=%d\n", len(os.Args))

	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer f.Close()

	// f, _ := os.Create("test.txt")
	// f.Write([]byte("hello\n"))
	// f.Write([]byte("hello2\n"))
	// f.Write([]byte("hello3\n"))
	// // f.Write([]byte("aaasasa\n"))
	// // f.WriteAt([]byte("Golang"), 1)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("Stringggggg")

	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	fmt.Println("bs:", bs)

	n, err := f.Read(bs) // fを読み込んでbsに書き込む
	fmt.Println(n)
	fmt.Println(string(n) + "string")

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 1)

	fmt.Println(nn)
	fmt.Println(string(bs2))

}
