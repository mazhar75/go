package main

import "fmt"

func main() {
	fmt.Printf("Please enter a number : ")
	var x int
	fmt.Scan(&x)
	if x&1 == 1 {
		fmt.Println("Your number is Odd")
	} else {
		fmt.Println("Your Number is even.")
	}

}
