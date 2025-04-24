package main

import "fmt"

func main(){
	var factorial = func (n int) int {
		res := 1
		for i:=2;i<=n;i++{
			res*=i
		}
		return res
	}
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}