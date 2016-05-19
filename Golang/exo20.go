package main

import (
	"strconv"
	"fmt"
	"math/big"
)

/*
	return the sum of digit
 */
func sumDigitInt(a *big.Int) int{
	sol := int(0)
	b := a.String()
	for _,j := range b {
		temp, _ := strconv.Atoi(string(j))
		sol += temp
	}
	return sol
}
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

func main() {
	n:= big.NewInt(100)
	fac := factorial(n)
	sol := sumDigitInt(fac)
	fmt.Println(sol)
}
