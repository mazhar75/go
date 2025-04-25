package main

import "fmt"

func process() func(x int, y int){
	return add
}

func add(a int,b int){
	c := a+b
	fmt.Println(c)
}

func main(){
	sum := process()
	sum(4,5)
}