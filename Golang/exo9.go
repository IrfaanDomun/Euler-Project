package main

import "fmt"

/*
	return true if a,b,c are a pythagorian triplet
 */
func isPytha(a,b,c int ) bool{
	if a*a+b*b ==c*c {
		return true
	}
	return false
}
/*
	return true if a+b+c = variable
 */
func triplet(variable ,a,b,c int) bool{
	if a+b+c == variable{
	return true
	}
	return false
}

/*
	return true  if its the right triplet
 */
func rightTriplet(variable,a,b,c int) bool {
	if isPytha(a,b,c){
		if triplet(variable,a,b,c){
			return true
		}
	}
	return false
}
/*
	find the solution
 */
func pythaTriplet(variable int)  int{

	for i:=1;i<variable;i++{
		for j:=1;j<variable;j++{
			for k:=1; k<variable;k++{
				if rightTriplet(variable,i,j,k){
					fmt.Println(i,j,k)
					return i*j*k
				}
			}
		}
	}
	return -1
}
func main() {
	variable := 1000
	sol:= pythaTriplet(variable)
	fmt.Println(sol)

}
