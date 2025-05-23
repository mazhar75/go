package main

import (
	"fmt"
	"sync"
	"time"
)
var arrTest = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var minicache sync.Map

func towrite(text string) {

	for _, v := range arrTest {
		minicache.Store(v, text)
		time.Sleep(10 * time.Millisecond)
	}

}

func main() {
	go towrite("test1")
	go towrite("test2")
	go func() {
		time.Sleep(10 * time.Millisecond)
		for _, v := range arrTest {
			if val, ok := minicache.Load(v); ok {
				fmt.Println(v, val)

			} else {
				fmt.Println(v, "not get")
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	time.Sleep(10 * time.Second)
}
