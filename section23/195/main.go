package main

import (
	"fmt"
	"os"
)


func CheckFiles (done <-chan interface{}, filenames ...string) <-chan *os.File{
	response := make(chan *os.File)

	go func(){
		defer close(response)
		for _,filename := range filenames {
			file, err := os.Open(filename)
			if err!= nil{
				fmt.Println(err)
				return
			}

			select{
				case <-done:
					return
				case response <- file:
			}
		}

	}()
	return response

}
func main() {
	done := make(chan interface{})

	defer close(done)

	filenames:= []string{"main.go","x.go"}

	for res:= range CheckFiles(done, filenames...){
		fmt.Printf("response: %v\n", res)
	}

}

// エラーハンドリングについての考え方
// エラーハンドリング誰がそのエラーを処理する責任を負うのかを考えることが大事

