package main

import "fmt"
func call(sl ...int)[]int{
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	return sl[2:4]
}

func main(){
	s := call(1,2,3,4,5)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}