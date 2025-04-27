package main
import "fmt"

func main(){
    var n int
	fmt.Print("Enter the the number you want to in slice : ")
	fmt.Scan(&n)

	var s []int
    var x int
	for i:=0;i<n;i++{
		fmt.Scan(&x)
		s=append(s,x)
		fmt.Printf("After %d iteration : \n",i+1)
		fmt.Println(s)
		fmt.Println(len(s))
		fmt.Println(cap(s))
  	}
}