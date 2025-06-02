/*
Post নামে একটি Struct তৈরি করুন যার ID, Category, Location, Time নামক ফিল্ড থাকবে। এখন Location এর মধ্যে আবার Latitude, Longitude থাকবে।
 Post এর দুটি ইন্সট্যান্স তৈরি করুন । শর্ত হচ্ছে, তাদের ID, Category, Time আলাদা কিন্তু Location একই থাকবে ।
*/
package main
import (
	"fmt"
	"time"
)
type Location struct{
	Latitude float64
	Longitude float64
}
type Post struct{
	ID string
	Category string
	Location Location
	Time time.Time
}


func main(){
    location := &Location{
		Latitude : 124.5,
		Longitude : 100.9,
	}
	post1 := &Post{
		ID : "9h82972jgf1",
		Category : "Fruits",
		Location : Location{Latitude : location.Latitude, Longitude : location.Longitude},
		Time : time.Now(),
	}
	//time.Sleep(30*time.Second)
	post2 := &Post{
		ID : "5htgdf5637",
		Category : "Fish",
		Location : Location{Latitude : location.Latitude, Longitude : location.Longitude},
		Time : time.Now(),
	}
	fmt.Println(post1)
	fmt.Println(post2)
}