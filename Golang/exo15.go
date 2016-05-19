package main

import (
	"fmt"
	"math/big"
)
/*
	return the factorial of n
 */
func factorial(n *big.Int) (result *big.Int) {
	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		// return n * factorial(n - 1);
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
}
/*
	compute the number of combination of path given for a movement of 'a' row and 'b' columns
 */
func combination(a,b int64) *big.Int {
	n:= big.NewInt(a)
	m := big.NewInt(b)
	o := big.NewInt(a-b)
	res := big.NewInt(0)
	res.Mul(factorial(m),factorial(o))
	res.Div(factorial(n),res)
	return res
}

func main() {
	n:= int64(20)
	m:=int64(20)
	fmt.Println(combination(n+m,n))
}
