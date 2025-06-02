/*
এই কোডটির আউটপুট কি হবে? বুঝার চেষ্টা করব আউটপুট কেন এমন হচ্ছে এবং বুঝতে পারলে ঠিক করে ফেলি –
*/
package main

import "fmt"

type Location struct {
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lon"`
}

type Post struct {
    Category string   `json:"category"`
    Time     int64    `json:"time"`
    Location Location `json:"location"`
}

func (post *Post) setLocation() {
    post.Location = Location{
        Latitude:  23.99,
        Longitude: 90.45,
    }
}

func main() {
    post := &Post{}
    post.setLocation()
    fmt.Println(post.Location)
}