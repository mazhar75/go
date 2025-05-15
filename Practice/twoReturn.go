//Statements : একটি ফাংশন লিখুন যেটি একটি সংখ্যা প্যারামিটার হিসেবে নিবে এবং সংখ্যাটি ধনাত্মক হলে সংখ্যাটির square root এবং positive মেসেজ return করবে , অন্যথায় -1 এবং negative মেসেজ return করবে –
package main

import "fmt"
func check(x int)(int,string){
	if x>0 {
		return x*x,"positive"
	} else {
		return -1,"negative"
	}
}
func main() {
    var n int
    fmt.Scan(&n)
    result, message := check(n)
    if message == "negative" {
        fmt.Println("this is a negative number, So square root of ", n, " is impossible")
    } else {
        fmt.Println("square root of ", n, " is ", result)
    }
}