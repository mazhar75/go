package main

import "fmt"
type Person struct {
	name string
	age int
	sex string
	adress string 
}
/*
func PrintDetails(user Person){

	fmt.Println("Person name   : ",user.name)
	fmt.Println("Person age    : ",user.age)
	fmt.Println("Person sex    : ",user.sex)
	fmt.Println("Person adress : ",user.adress)
}
*/
func (user Person) PrintDetails(){
	fmt.Println("Person name   : ",user.name)
	fmt.Println("Person age    : ",user.age)
	fmt.Println("Person sex    : ",user.sex)
	fmt.Println("Person adress : ",user.adress)
}
func main(){
    person1 := Person {
		name   : "Md. Mazharul Islam",
		age    : 25,
		sex    : "Male",
		adress : "Shibganj, Sylhet",
	}
	person1.PrintDetails()

	var userList [5]Person
	userList[0].name="Arham"
	userList[0].age=25
	userList[0].sex="male"
	userList[0].adress="Shahi Eidgah, Sylhet"
	userList[0].PrintDetails()
	


}