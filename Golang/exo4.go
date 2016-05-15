package main

import (
	"fmt"
	str "strconv"
)
/*
	reverse a string
 */
func reverse( s string) (res string){

	for _,j := range s {
		res = string(j)+res
	}
	return res
}
/*
	Is a Palindrome
 */
func isPalindrome(s string) bool{
	if s == reverse(s){
		return true
	}
	return false
}

/*
	biggest palindrome product of two 3 digit numbers
 */
func biggestPalindrome() int{
	max:=0
	for i:=100;i<1000;i++{
		for j:=100;j<1000;j++{
			N:=i*j
			n:= str.Itoa(N)
			if isPalindrome(n){
				if max < N{
					max = N
				}
			}
		}
	}
	return max
}
/*
	Get the result
 */
func main() {
	fmt.Println(biggestPalindrome())
}
