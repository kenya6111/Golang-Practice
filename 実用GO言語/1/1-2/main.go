package one_2

import "fmt"

type ErrDatabase int

//	func (e ErrDatabase) Error() string {
//		return "err"
//	}
// func (e ErrDatabase) String() string {
// 	return "stringer"
// }

const (
	errDatabase ErrDatabase = 0
	aaa         ErrDatabase = 1
	vdfd        ErrDatabase = 3
)

func TestOne_2() {
	fmt.Println(errDatabase)
	fmt.Println(aaa)
	fmt.Println(vdfd)
}
