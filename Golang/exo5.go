package main

import (
	"fmt"
)

/*
	get the smallest multiple of 1-variable
 */
func smallestMultiple(variable int) int{
	i:=1
	for {
		isRes := true
		for j:=1;j<=variable;j++{
			if i%j!=0{
				isRes = false
				break
			}

		}
		if isRes{
			return i
		}
		i++

	}
}

func main() {
	variable :=20
	fmt.Println(smallestMultiple(variable))
}
