/*
Problem Statement:একটি প্রোগ্রাম লিখুন যা দুটি গো-রুটিন তৈরি করে। প্রথম গো-রুটিনে 0 থেকে 10 পর্যন্ত ইন্টিজার তৈরি করে এবং একটি Channel এর মাধ্যমে দ্বিতীয় গো-রুটিনে পাঠায়। 
দ্বিতীয় গো-রুটিন ইন্টিজার গ্রহণ করে এবং সেগুলোকে কনসোলে প্রিন্ট করে।
*/
package main
import (
	"fmt"
)

func main(){

	ch := make(chan int)
	go func() {
		for val := range ch {
			fmt.Println("Received via:", val)
		}
	}()
    for i := 0; i <= 10; i++ {
       ch <- i
    }
    defer close(ch)

}