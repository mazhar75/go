/*
Employee  নামের একটি Struct তৈরি করুন যার name, age এবং salary ফিল্ড থাকবে। 
GiveRaise নামক একটি মেথড তৈরি করুন যা একজন কর্মচারীর বেতন একটি নির্দিষ্ট শতাংশ হারে বৃদ্ধি করে। 
এমন একটি প্রোগ্রাম লিখুন যা একটি Employee ইন্সট্যান্স তৈরি করে, GiveRaise মেথড এর সাহায্যে বেতন  বৃদ্ধি করে এবং তাদের নতুন বেতন প্রিন্ট করে।
*/
package main
import "fmt"

type Method interface{
	GiveRaise(percentage float64)
}
type Employee struct{
	name string `json:"name"`
	age int32   `json:"age"`
	salary float64 `json:"salary"`
}
func (employee *Employee) GiveRaise(percentage float64){
	employee.salary=employee.salary+percentage*employee.salary/100
}
func main(){
	employee := &Employee{
		name   : "Mazhar",
		age    : 25,
		salary : 60000,
	}
	fmt.Println(employee)
	employee.GiveRaise(10)
	fmt.Println(employee)
}