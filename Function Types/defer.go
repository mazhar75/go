package main

import "fmt"

func test() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    fmt.Println("Function Body")
}


func main(){
	test()
}