/*
গো-রুটিন  এবং চ্যানেল ব্যবহার করে জোড় এবং বিজোড় সংখ্যা বের করার একটি কোড লিখুন।
*/
package main

import "fmt"

func Checker(ch chan string,x int){
	if x%2==0{
		ch <-"Even"
	} else {
		ch <-"Odd"
	}
}
func main(){
	ch := make(chan string)
	var x int
	fmt.Scan(&x)
	go Checker(ch,x)
	fmt.Println(<-ch)
}