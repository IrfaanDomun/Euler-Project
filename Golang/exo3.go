package main
import (
	"fmt"
	"math"
)
/*
	Tell if i is multiple of variable
 */
func isMultiple(variable int64,i int64) bool{
	var a int64 =variable/i
	if variable == a*i{
		return true
	}
	return false
}
/*
	Tell if i is prime
 */
func isPrime(i int64) bool{
	var j int64 = 2
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
	Get the longest prime
 */

func longestPrime(variable int64)int64{
	var sol int64= 1
	var i int64 = 1
	for float64(i)<=math.Sqrt(float64(variable)){
		if isMultiple(variable,i){
			if isPrime(i) {
				sol = i
			}
		}
		i++
	}
	return sol
}

func main() {
	var variable int64 =600851475143
	fmt.Println(longestPrime(variable))
}
