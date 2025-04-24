package main

import "fmt"

func main(){
	
    func (n int){
		res := 1
		for i:=2;i<=n;i++{
			res*=i
		}
		fmt.Println(res)
	}(5)

	
}