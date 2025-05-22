package main
import (
	"fmt"
	"time"
)
func IntegerLoop(){
	for i:=1;i<=5;i++{
		time.Sleep(1*time.Second)
		fmt.Printf("Integer Loop : %d\n",i)
	}
}
func CharacterLoop(){
	for c:='a';c<='e';c++{
		time.Sleep(2*time.Second)
		fmt.Printf("Character Loop : %c\n",c)
	}
}
func main(){
	fmt.Println("Main go routine Started!")
	go IntegerLoop();
	go CharacterLoop();
	time.Sleep(12*time.Second)
	fmt.Println("Main go routine Ended!")
}
