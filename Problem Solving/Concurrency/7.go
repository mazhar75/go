/*
গো-রুটিন এবং চ্যানেল ব্যবহার করে ১ থেকে n পর্যন্ত সংখ্যার বর্গের যোগফল বের করার একটি কোড লিখুন.
*/
package main

import (
	"fmt"
	"time"
)

func SumOfSquares(c, quit chan int) {
    for i := 1; i <= 6; i++ {
        c <- i * i 
    }
    quit <- 0 
}

func main() {
    mychannel := make(chan int)
    quitchannel := make(chan int)
    sum := 0

    go func() {
        for {
            select {
            case val := <-mychannel:
                sum += val
            case <-quitchannel:
                fmt.Println("Sum of squares:", sum)
                return
            }
        }
    }()

    SumOfSquares(mychannel, quitchannel)
	time.Sleep(1*time.Second)

}
