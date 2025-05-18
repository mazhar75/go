//main.go
package main
import (
	"fmt"
)

func main(){
    

      user1:=&User{
		Gmail:"mazharulnihadmd@gmail.com",
		Name:"Mazharul Islam",
		Age : 25,
		Sex : "Male",
		Address:Adress{
           Village:"Uchail",
		   Postcode:3300,
		   District:"Habiganj",
		   Division:"Sylhet",
		},
	  }
	  fmt.Println("Account 1 Operation")
	  fmt.Println(user1)
	  user1.AccountList=append(user1.AccountList,Account{"7fght423fsg",0.0})
	  fmt.Println(user1)
	  user1.AccountList=append(user1.AccountList,Account{"hss765wfsgs",500.0})
	  fmt.Println(user1)
	  fmt.Println(user1.AccountList[0].CheckBalance())
	  msg,err:=user1.AccountList[0].Deposit(509.0)
	  if err!=nil{
		fmt.Println(err)
	  } else {
		fmt.Println(msg)
	  }
	  fmt.Println(user1.AccountList[0].CheckBalance())
	  msg,err=user1.AccountList[0].Withdraw(50.0)
	  if err!=nil{
		fmt.Println(err)
	  } else {
		fmt.Println(msg)
	  }
	  fmt.Println(user1.AccountList[0].CheckBalance())
	  fmt.Println("---------------------------------")
	  fmt.Println("Account 2 Operation")
	  fmt.Println(user1.AccountList[1].CheckBalance())
	  msg,err=user1.AccountList[1].Deposit(50.0)
	  if err!=nil{
		fmt.Println(err)
	  } else {
		fmt.Println(msg)
	  }
	  msg,err=user1.AccountList[1].Withdraw(5000.0)
	  if err!=nil{
		fmt.Println(err)
	  } else {
		fmt.Println(msg)
	  }





}