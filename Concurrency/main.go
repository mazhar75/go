package main

import (
  "fmt"
  "time"
)

func Hello(){
	fmt.Println("Hi , I am Mazhar !")
}
func Hi(){
	fmt.Println("Hi , I am Nihad !")
}
func init(){
    go Hello()
	time.Sleep(1*time.Second)

}
func main(){
	go Hi()
	time.Sleep(1*time.Second)
}