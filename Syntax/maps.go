package main 

import "fmt"

func main(){

	var mp map[string]int = make(map[string]int)
	mp["nihad"]=75
	mp["arham"]=5
	for name,roll := range mp{
		fmt.Println(name,roll)
	}
}