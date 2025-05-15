// Statements : একটি ফাংশন লিখুন, যা একটি ইন্টিজারের স্লাইস রিসিভ করবে এবং এর সংখ্যাগুলোর গড় রিটার্ন করবে।
package main

import "fmt"
func average(numbers []int) float64{
	 sum :=0
	 for _,x:=range numbers{
		sum+=x
	 }
	 avg:=float64(sum)/float64(len(numbers))
	 return avg
}
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    avg := average(numbers)
    fmt.Printf("The average of %v is %.2f\n", numbers, avg)
}