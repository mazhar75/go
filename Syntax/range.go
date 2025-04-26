package main 

import "fmt"

func main(){
   var s []int //nil slice
   s=append(s,2)
   s=append(s,3)
   s=append(s,10)

   // range loop

   for i,x := range s{
	 fmt.Println(i,x)
   }
}