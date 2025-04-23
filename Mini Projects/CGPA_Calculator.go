//This is my first project in go
package main
import "fmt"
func main(){
	fmt.Println("Please provide necessary info to calculate your cgpa.")
	fmt.Print("Enter your current cgpa & completed credits : ")
	var curCGPA float32
	var curTotalCredit float32
	fmt.Scan(&curCGPA,&curTotalCredit)
	var totalSub int
	fmt.Print("Enter total number subject in this semester : ")
	fmt.Scan(&totalSub)
	var totalPoint,totalCredit,subGPA,subCredit float32

	for i:=1;i<=totalSub;i++{
		fmt.Printf("For subject %d enter GPA & Course Credit: ",i)
		fmt.Scan(&subGPA,&subCredit)
		totalPoint+=subGPA*subCredit
		totalCredit+=subCredit
	}
	cgpa:=totalPoint/totalCredit
	fmt.Printf("This semester your cgpa is %.2f\n",cgpa)
	curCGPA=((curCGPA*curTotalCredit)+(cgpa*totalCredit))/(curTotalCredit+totalCredit)
	fmt.Printf("Your total cgpa till now is %.2f\n",curCGPA)

	
}