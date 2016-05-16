package  main

import (
	"math"
	"fmt"
)
/*
	Tell if i is prime
 */
func isPrime(i int) bool{
	var j int = 2
	lim := float64(i)
	for j<=int(math.Sqrt(lim)){
		if i%j==0{
			return false
		}
		j++
	}
	return true
}

/*
	sum of n prime number
 */
func sumPrime(n int) int{
	sum :=0
	i := 2
	for {

		if i > n{
			return sum
		}
		if isPrime(i){
			sum+=i
		}
		i++
	}
}
func main() {
	variable := 2000000
	fmt.Println(sumPrime(variable))