package main

import (
    "fmt"
)

func producer(chnl chan int) {
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}

func main() {
    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
// Output:
// Received  0 true
// Received  1 true
// Received  2 true
// Received  3 true
// Received  4 true
// Received  5 true
// Received  6 true
// Received  7 true
// Received  8 true
// Received  9 true