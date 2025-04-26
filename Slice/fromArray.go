package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5,6}


	//slice
	s1 := a[1:5]
	s2 := s1[2:4]
	s2[0]=11
	s2 = append(s2,-1)
	s2 = append(s2,-2)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(cap(a))
}