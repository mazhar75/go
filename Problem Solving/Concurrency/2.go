/*
একটি প্রোগ্রাম লিখুন যা কনসোল থেকে পূর্ণসংখ্যা পড়ে এবং একটি চ্যানেলের মাধ্যমে একটি গো-রুটিনে পাঠায়। গো-রুটিন কনসোলে প্রতিটি পূর্ণসংখ্যার বর্গ প্রিন্ট করে।
*/


package main

import (
	"fmt"
	"time"
)

func Print(ch <-chan int) {
	for val := range ch {
		if val == -1 {
			fmt.Println("Exiting...")
			return
		}
		fmt.Println("Square of your number:", val*val)
	}
}

func main() {
	ch := make(chan int)

	// Start single receiver goroutine
	go Print(ch)

	for {
		fmt.Printf("Input your number (or -1 to exit): ")
		var x int
		fmt.Scan(&x)
		ch <- x
		if x == -1 {
            time.Sleep(2*time.Second)
			break
		}
	}

	close(ch) // Close channel after sending done
}
