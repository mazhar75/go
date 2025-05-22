package main
import (
	"fmt"
	"time"
)
func Worker(cancel <-chan struct{}){
    for{
		select{
		  case <-cancel:
			 fmt.Println("Got signal to cancel !")
			 return
		  default :
		      time.Sleep(1*time.Second)
		      fmt.Println("Working")
			  	 
		}
	}
}
func main(){
	fmt.Println("Main go routine Started ")
	cancel := make(chan struct{})
	fmt.Println("Worker go routine Started ")
	go Worker(cancel)  // pass the channel to worker
	time.Sleep(3*time.Second)
	fmt.Println("Worker go routine Cancel signal ")
	close(cancel)  // send stop signal
	time.Sleep(1*time.Second)
	fmt.Println("Main go routine Ended ! ")


}