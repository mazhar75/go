package mod
import "errors"

func Mod(num1 int, num2 int) (int,error){

	if(num2==0){
        return -1,errors.New("Cann't divide by zero")
	} 
	return num1/num2,nil
}