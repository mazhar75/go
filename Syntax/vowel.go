package main

import "fmt"

func main() {
	var ch string
	fmt.Print("Enter your character (only Lowercase): ")
	fmt.Scan(&ch)
	var x = rune(ch[0])

	var found bool = false
	switch x {
	case 'a', 'e', 'i', 'o', 'u':
		found = true
	default:
		found = false
	}
	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
