package main

import (
	"fmt"
	"math"
)

/*
	return the nth triangle number
 */
func triangleNumber(n int) int{
	sol := 0
	for i:=0;i<=n;i++{
		sol+=i
	}
	return sol
}

/*
	return the number of divisor
 */
func nbDivisor(n int) int {
	res :=0
	for i:=1;i<int(math.Sqrt(float64(n))) +1;i++{
		if n%i == 0{
			res+=2
		}
	}
	return res
}
/*
	return the first triangle number with above n divisor
 */
func firstDivisorTriangle(n int ) int{
	i:=1
	nb :=0
	for {
		nb += i
		if nbDivisor(nb) >=n{

			return nb
		}
		i++
	}
}
func main() {
	n:=500
	sol := firstDivisorTriangle(n)
	fmt.Println(sol)
}
