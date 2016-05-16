package main

import "fmt"

/*
	compute the Collatz sequence
 */
func collatz(n int) int{
	if n%2==0{
		return n/2
	}
	return 3*n+1
}

/*
	find the longest chain under variable
 */
func longestChain(variable int) int {
	max :=0
	sol :=0
	for i:=1;i<variable;i++{
		temp:=i
		count :=1
		for {
			temp = collatz(temp)
			count+=1
			if temp ==1{
				break
			}
		}
		if count > max {
			max = count
			sol = i
		}
	}
	return sol
}
func main() {
	variable := 1000000
	res := longestChain(variable)
	fmt.Println(res)
}
