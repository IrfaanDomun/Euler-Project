package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
	return the sum of digit
 */
func sumDigit(a float64) int{
	sol := int(0)
	b := strconv.FormatFloat(a,'f',0,64)
	for _,j := range b {
		temp, _ := strconv.Atoi(string(j))
		sol += temp
	}
	return sol
}
func main() {
	a := math.Pow(2,1000)
	sol := sumDigit(a)
	fmt.Println(sol)
}
