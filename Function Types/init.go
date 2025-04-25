package main

import "fmt"

//This function will executed after main
func main(){
	fmt.Println("This is main function")
}

//This function will be executed first
func init(){
	fmt.Println("This is init function")
}