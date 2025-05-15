//statement : একটি ফাংশন লিখুন যেটি বিভিন্ন সংখ্যক একই ডাটা টাইপের ইনপুট আর্গুমেন্ট হিসেবে গ্রহণ করতে পারবে , সংখ্যা গুলো যোগ করে যোগফল return করবে – 

package main
import "fmt"

func sumAll(ar ...int)int{
	sum:=0
	for _,val:=range ar{
		sum+=val
	}
	return sum
}
func main() {
    arr1 := []int{1, 2, 3}
    arr2 := []int{1, 2, 3, 4}
    arr3 := []int{1, 2, 3, 4, 5}
    result1 := sumAll(arr1...)
    result2 := sumAll(arr2...)
    result3 := sumAll(arr3...)
    fmt.Println("sum of first array ", result1)
    fmt.Println("sum of second array ", result2)
    fmt.Println("sum of third array ", result3)
}