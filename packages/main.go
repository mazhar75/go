package main

import (
	"fmt"
	"example.com/add"
	"example.com/subtract"
	"example.com/multiply"
	"example.com/divide"
	"example.com/mod"

)
func main(){
	var a,b int
	fmt.Scan(&a,&b)
	sum:=add.Add(a,b)
	sub:=subtract.Subtract(a,b)
	mul:=multiply.Multiply(a,b)
    div,er1:=divide.Divide(a,b)
	mod,er2:=mod.Mod(a,b)
	
	fmt.Println("Sum:", sum)
	fmt.Println("Subtract:", sub)
	fmt.Println("Multiply:", mul)
	if(er1==nil){
		fmt.Println("Divide:", div)
	} else {
		fmt.Println(er1)
	}

	if(er2==nil){
		fmt.Println("Modulus:", mod)
	} else {
		fmt.Println(er2)
	}
	



}