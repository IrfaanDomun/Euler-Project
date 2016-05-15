package main

import (
	"fmt"
)

func fibo(variable int64) int64{
	var f int64 = 1
	var s int64 = 2
	var sol int64 = 2
	var fib int64 = 0
	for fib <= variable{
		fib=f+s
		if fib%2 ==0{
			sol+=fib
		}
		f = s
		s =fib
		//fmt.Println(fib)
	}
	return sol
}

func main() {
	var variable int64= 4000000
	fmt.Println(fibo(variable))
}
