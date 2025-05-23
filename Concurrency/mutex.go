package main

import (
    "fmt"
    "sync"
)
var counter int
var mutex sync.Mutex

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go increment(&wg,i)
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}

func increment(wg *sync.WaitGroup,x int) {
	fmt.Printf("Before Blocking goroutine %d\n",x)
    mutex.Lock()
    defer mutex.Unlock()
    defer wg.Done()
	fmt.Printf("After Blocking goroutine %d\n",x)
    counter++
}

// Output :
// Counter: 10