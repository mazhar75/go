/*
PrintValues নামে একটি ফাংশন লিখুন যা এম্পটি ইন্টারফেসের একটি স্লাইস নেয় এবং স্লাইসের প্রতিটি উপাদানের মান প্রিন্ট করে। 
আপনার ফাংশন বিভিন্ন ধরনের মান হ্যান্ডেল করতে সক্ষম হতে হবে।
*/
package main

import (
	"fmt"
)

func PrintValues(values []interface{}) {
	for i, v := range values {
		switch v := v.(type) {
		case int:
			fmt.Printf("Index %d: int value = %d\n", i, v)
		case float64:
			fmt.Printf("Index %d: float64 value = %f\n", i, v)
		case string:
			fmt.Printf("Index %d: string value = %s\n", i, v)
		case bool:
			fmt.Printf("Index %d: bool value = %t\n", i, v)
		default:
			fmt.Printf("Index %d: unknown type (%T) value = %v\n", i, v, v)
		}
	}
}

func main() {
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}}
	PrintValues(values)
}
