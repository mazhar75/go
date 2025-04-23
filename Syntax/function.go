package main

import "fmt"

func Cal(num1 int, num2 int) (int,int){
	return num1+num2,num1*num2
}
func main(){
	var a,b int
	fmt.Scan(&a,&b)
	sum,mul:= Cal(a,b)
	fmt.Printf("Sum & Multiplication of two numbers %d and %d is: %d & %d\n", a, b, sum,mul)

}