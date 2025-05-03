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
	jsonString := `{"name":"Mazhar","age":22,"email":"mazhar@example.com"}`

	var p Person

	// Convert JSON to struct
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Printf("Struct: %+v\n", p)
}
