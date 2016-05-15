package main

import (
	"fmt"
)
func sumMul(variable int) int {
	sol :=0
	for mul3:=0; mul3 < variable;mul3+= 3{
		if mul3%5 != 0 {
			sol += mul3
		}
	}
	for mul5:=0; mul5 < variable ;mul5+=5{
		sol+=mul5
	}
	return sol
}

func main(){
	variable := 1000
	fmt.Println(sumMul(variable))
}