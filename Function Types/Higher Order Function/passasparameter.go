package main

import "fmt"

func process(a int, b int, op func (x int, y int)){
	  op(a,b)
}
func add(a int, b int){
	c := a+b
	fmt.Println("Sum of two numbers is : ",c)
}
func sub(a int, b int){
	 c := a-b
	 fmt.Println("Subtraction of two numbers is : ",c) 
}

func multiply(a int, b int){
	 c := a*b
	 fmt.Println("Multiplication of two numbers is : ",c)
}

func divide(a int, b int){
	if b==0{
		fmt.Println("Cann't divide by zero")
		return
	}
	c := a/b
	fmt.Println("Division of two numbers is : ",c)
}

func mod(a int, b int){
	if b==0 {
		fmt.Println("Cann't divide by zero")
		return
	}
	c := a%b
	fmt.Println("Modular of two numbers is : ",c)
}

func main(){
	process(8,5,add)
	process(8,5,sub)
	process(8,5,multiply)
	process(8,0,divide)
	process(8,5,mod)

}