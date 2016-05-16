package main

import ( "math"
"fmt"
)
/*
	Tell if i is multiple of variable
 */
func isMultiple(variable int,i int) bool {
	var a int = variable / i
	if variable == a * i {
		return true
	}
	return false
}

/*
	Tell if i is prime
 */
func isPrime(i int) bool{
	var j int = 2
	lim := float64(i)
	for float64(j)<=math.Sqrt(lim){
		if isMultiple(i,j){
			return false
		}
		j++
	}
	return true
}
/*
	find the n prime number
 */
func nPrime(variable int) int {
	n:=1
	i:=3
	for{
		if isPrime(i){
			n+=1
			if n==variable{
				return i
			}
		}
		i+=2
	}
}
func main() {
	variable := 10001
	fmt.Println(nPrime(variable))
}
