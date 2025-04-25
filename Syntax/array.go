package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	// array size must be at compile time else error
	var a [5]int
	for i:=0;i<5;i++{
		fmt.Scan(&a[i])
	}
	for i:=0;i<5;i++{
		fmt.Print(a[i]," ")
	}
	fmt.Println(" ")
}