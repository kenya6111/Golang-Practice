package main

import "fmt"

type MyError struct {
	message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.message
}

func RaiseError() error {
	return &MyError{message: "カスタムエラーが発生しました", ErrCode: 1234}
}
func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e := MyError{message: "errメッセージ", ErrCode: 4444}
	fmt.Println(e)
	fmt.Println(e.Error())

}
