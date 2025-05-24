/*
একটি প্রোগ্রাম লিখুন যেখানে একটি ইন্টিজার Buffered চ্যানেল তৈরি করতে হবে, যার বাফার সাইজ হবে 5। প্রোগ্রামটি দুটি গো-রুটিন তৈরি করবে যা চ্যানেলে এবং চ্যানেল থেকে ক্রমাগত ইন্টিজার পাঠাবে এবং গ্রহণ করবে। 
সেন্ডার গো-রুটিনে 1 থেকে 10 পর্যন্ত পূর্ণসংখ্যা পাঠাতে হবে। যখন রিসিভার গো-রুটিন চ্যানেল থেকে রিসিভ করবে, তখন প্রতিটি সংখ্যা  প্রিন্ট করবে।
*/
package main
import (
	"fmt"
	"time"
)

func Buffer(ch <-chan int){
	for val:=range ch{
		fmt.Println("Recieved Value : ",val)
	}
}
func main(){
	ch:=make(chan int,5)
	go Buffer(ch)
	for i:=0;i<10;i++{
		ch<-i
	}
	time.Sleep(2*time.Second)
	close(ch)
}