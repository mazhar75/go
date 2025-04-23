package main

import "fmt"

func main() {
	var x int32 = 56
	var y float64 = 76.09
	var z string = "Hi, I am Mazhar"
	var b bool = true
	var ch rune = 'I'
	fmt.Println(x, y, z, b, ch)
	//for pointer
	var p *int32
	p = &x
	fmt.Println(*p)

	//User Input
	fmt.Scan(&x, &y, &z, &b, &ch)
	fmt.Println("Your Inputed ouput:", x, y, z, b, ch)

}
