package main

import (
	"fmt"
)

/*
	return the sum of square
 */
func sumSquare(variable int) int{
	res:=0
	for i:=0;i<=variable;i++{
		res += i*i
	}
	return res
}
/*
	return the square of sum
 */
func squareSum(variable int) int{
	res:=0
	for i:=0;i<=variable;i++{
		res += i
	}
	return res*res
}

/*
	return the difference of the two
 */
func diffSquare(variable int) int{
	return squareSum(variable)-sumSquare(variable)
}
/*
	print the result
 */
func main() {
	variable := 100
	fmt.Println(diffSquare(variable))
}