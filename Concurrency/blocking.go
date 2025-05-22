package main
import (
	"fmt"
	"time"
)

func test_blocking(done chan bool){
	for i:=1;i<=5;i++{
		time.Sleep(1*time.Second)
		fmt.Println("Hello World !")
	}
	done <- true
}
func main(){
	fmt.Println("Main go routine started ")
	fmt.Println("test_bloking go routine ")
	done := make(chan bool)
	go test_blocking(done)
	fmt.Println("waiting for recieve !")
	<-done
	fmt.Println("Tata bye bye khotom geya!")
}

