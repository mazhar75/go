package main

import "fmt"

func main() {

	var t int
	fmt.Scan(&t)
    k:=0
	for k<t{
	   var n,x int
	   fmt.Scan(&n,&x)
	   for i:=0;i<n;i++{
		  if i==x{
			continue;
		  } else {
			 fmt.Print(i," ")
		  }
	   }
	   if x<n{
	    fmt.Println(x)
	   }
	   k++
	}
}
