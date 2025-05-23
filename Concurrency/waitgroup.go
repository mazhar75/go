package main

import (
	"fmt"
	"time"
	"sync"
)

func Go_Routine_1(wg *sync.WaitGroup, x int){
	defer wg.Done()
	fmt.Printf("Started Routine %d in Go_Routine_1\n",x)
	time.Sleep(1*time.Second)
	fmt.Printf("Finished Routine %d in Go_Routine_1\n",x)
}
func Go_Routine_2(wg *sync.WaitGroup,x int){
	defer wg.Done()
	fmt.Printf("Started Routine %d in Go_Routine_2\n",x)
	time.Sleep(2*time.Second)
	fmt.Printf("Finished Routine %d in Go_Routine_2\n",x)
}
func main(){
	fmt.Println("Starting main Goroutine")
    var waitgroup sync.WaitGroup
	for i:=1;i<=5;i++{
		if i%2==1{
			waitgroup.Add(1)
			go Go_Routine_1(&waitgroup,i)
		} else {
			waitgroup.Add(1)
			go Go_Routine_2(&waitgroup,i)
		}
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}