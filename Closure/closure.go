package main

import "fmt"
func closure() func (a int) int {
     b := 5
	 return func (a int) int {
		b++
		return a+b
	 }
}

func main(){
    access := closure()
	sum := access(9)
	fmt.Println(sum)
	sum = access(9)
	fmt.Println(sum)

}

func init(){
	fmt.Println("Hi , This is Mazhar!")
}