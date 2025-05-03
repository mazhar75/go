package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p := Person{
		Name:  "Mazhar",
		Age:   22,
		Email: "mazhar@example.com",
	}

	// Marshal with indentation
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
