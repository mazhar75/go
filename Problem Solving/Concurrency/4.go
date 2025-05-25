/*
এমন একটি প্রোগ্রাম তৈরি করতে হবে যা একইসময়ে ইন্টারনেট থেকে গো-রুটিন ব্যবহার করে একাধিক ফাইল ডাউনলোড করবে। নিশ্চিত করতে হবে যে প্রোগ্রামটি বন্ধ করার আগে সমস্ত ডাউনলোড সম্পূর্ণ হয়েছে।
 sync.WaitGroup প্যাকেজ ব্যবহার করে এটা করতে হবে।আপনি প্রতিটি ফাইল ডাউনলোডের জন্য একটি নতুন গো-রুটিন তৈরি করবেন এবং সক্রিয় গো-রুটিন সংখ্যার ট্র্যাক রাখতে একটি ওয়েটগ্রুপ ব্যবহার করবেন৷ 
*/
package main

import(
	"fmt"
	"sync"
	"time"
)
func download(x int, wg *sync.WaitGroup){
	defer wg.Done()
    fmt.Printf("File downloading started at routine-%d\n",x)
	fmt.Printf("Downloading at routine %d ......\n",x)
	time.Sleep(1*time.Second)
	fmt.Printf("File downloading finished at routine-%d\n",x)
}
func main(){
	var wg sync.WaitGroup
	fmt.Println("Enter the number of files want to download : ")
	var n int
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		wg.Add(1)
		go download(i,&wg)
	}
	wg.Wait()
}