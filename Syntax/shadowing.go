package main
import "fmt"
var x=10
func main(){

   fmt.Println("Before block x's value : ",x)
   trick := true
   if trick {
	  x:=67
	  fmt.Println("x is the shadow of global variable")
	  fmt.Println("x's value here is :",x)
   }
   fmt.Println("After block x's value : ",x)
}