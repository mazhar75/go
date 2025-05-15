//Statement : একটি ফাংশন লিখুন যেটি একটি ফাইল create করবে এবং ফাইলটিতে “Hello World” লিখার পর defer ব্যবহার করে ফাইলটি close করে দিবে –
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    // Create a file with file name example.txt
    file, err := os.Create("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // 👈 Ensures file is closed when function ends

    // Write something to example.txt
    _, err = file.WriteString("Hello, World!")
    if err != nil {
        log.Fatal(err)
    }

    // Read from example.txt
    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(content))
}
