package main

import (
	"fmt"
	"time"
)

func BufferedChannel(ch chan time.Time) {
	for {
		t, ok := <-ch
		if ok {
			fmt.Println(t)
		} else {
			break
		}
	}
}

func main() {
	ch := make(chan time.Time, 5)

	// Launch reader goroutine
	go BufferedChannel(ch)

	// Fill the buffered channel
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second) 
		ch <- time.Now()
	}

	close(ch) // Close channel to signal end
	time.Sleep(2 * time.Second) 
}
