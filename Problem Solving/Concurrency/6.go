/*
একটি Go প্রোগ্রাম লিখুন যাতে ২ টি গো-রুটিন কল করা হবে। একটিতে ০ থেকে ১০০ পর্যন্ত জোড় সংখ্যা ও অন্যটিতে বিজোড় সংখ্যা, নির্দিষ্ট চ্যানেলে পাঠানো হবে। 
কন্ডিশনালের মাধ্যমে ০ থেকে ১০০ ক্রমানুসারে প্রিন্ট করতে হবে। ক্রম ঠিক রাখতে চ্যানেল ব্যবহার করতে হবে।
*/
package main
import (
	"fmt"
	"sync"
)


func printEven(evenCh chan bool, controlCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <=100; i += 2 {
		<-controlCh 
		fmt.Println(i)
		if i<100 {
			evenCh <- true 
		}		
	}
}

func printOdd(oddCh chan bool, controlCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 100; i += 2 {
		<-controlCh
		fmt.Println(i)
		oddCh <- true 
	}
}

func main() {
	evenCh := make(chan bool)
	oddCh := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	go printEven(evenCh, oddCh, &wg)
	go printOdd(oddCh, evenCh, &wg)
	oddCh <- true

	wg.Wait()
}




/*
//My solution : ok but not synchronized
func PrintEven(ch <-chan int, wg *sync.WaitGroup){
     defer wg.Done()
	 fmt.Println(<-ch)
}
func PrintOdd(ch <-chan int, wg *sync.WaitGroup){
     defer wg.Done()
	 fmt.Println(<-ch)
}

func main(){
	Oddch := make(chan int)
	Evench :=make(chan int)
	var wg sync.WaitGroup
	defer close(Oddch)
	defer close(Evench)
	for i:=0;i<100;i++{
		if i%2==0{
			wg.Add(1)
			go PrintEven(Evench,&wg)
			Evench <- i
		} else {
			wg.Add(1)
			go PrintOdd(Oddch,&wg)
			Oddch <- i
		}
	}
	wg.Wait()
}
*/
