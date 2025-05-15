//Question : একটি ফাংশন লিখুন যেটি দুটি সংখ্যা প্যারামিটার হিসেবে নেয় , সংখ্যা দুটি যোগ করে এবং যোগফল return করে – 

package main

import "fmt"

func sum(a int, b int) int {
	return a+b
}

func main(){
    x, y := 10, 20
    result := sum(x, y)
    fmt.Printf("The sum of %d and %d is %d\n", x, y, result)
}