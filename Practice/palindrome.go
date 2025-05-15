// Statements : একটি ফাংশন লিখুন, যা একটি String রিসিভ করবে এবং একটি Boolean রিটার্ন করবে। যদি String এর সকল ক্যারেক্টার ব্যবহার করে পেলিনড্রম বানানো সম্ভব হয় তাহলে true, আর নয়তো false রিটার্ন করবে।
package main
import "fmt"
func isPalindromePossible(str string) bool{
	 i:=0
	 j:=len(str)-1
	 for i<j{
		if(str[i]!=str[j]){
			return false
		}
		i++
		j--
	 }
	 return true
}
func main() {
    str := "madam"
    if isPalindromePossible(str) {
        fmt.Printf("%q is a palindrome\n", str)
    } else {
        fmt.Printf("%q is not a palindrome\n", str)
    }
}