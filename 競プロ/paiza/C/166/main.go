package q166

import (
	"flag"
	"fmt"
	"strconv"
)


func Main(){
	flag.Parse()
	num := flag.Arg(0)
	newnum,err := strconv.Atoi(num)

	if err != nil{
		return
	}


	a := getCoinNums(newnum)

	fmt.Println(a)


}


func getCoinNums (x int) int{
	a := x/500
	x -= a*500
	fmt.Println("a:",a)
	
	b := x/100
	x -= b*100
	fmt.Println("b:",b)
	
	c := x/50
	x -= c*50
	fmt.Println("c:",c)
	
	d := x/10
	x -= d*10
	fmt.Println("d:",d)
	
	e := x/5
	x -= e*5
	fmt.Println("e:",e)

	return a + b + c + d + e;

}