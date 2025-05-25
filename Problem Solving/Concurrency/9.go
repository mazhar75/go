/*
গো-রুটিন ব্যবহার করে [9,4,3,6,1,2,10,5,7,8] অ্যারেটি সর্টিং এর জন্য একটি কোড লিখুন।
*/
package main

import (
    "fmt"
)

// দুইটা সর্টেড স্লাইস ম্যার্জ করে দিবে
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    // বাকি অংশ যোগ করা
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}

// গো-রুটিন ব্যবহার করে মের্জ সর্ট
func concurrentMergeSort(arr []int, ch chan []int) {
    if len(arr) <= 1 {
        ch <- arr
        return
    }

    mid := len(arr) / 2

    leftCh := make(chan []int)
    rightCh := make(chan []int)

    // বাম পাশে গো-রুটিনে সর্ট করা
    go concurrentMergeSort(arr[:mid], leftCh)
    // ডান পাশে গো-রুটিনে সর্ট করা
    go concurrentMergeSort(arr[mid:], rightCh)

    left := <-leftCh
    right := <-rightCh

    ch <- merge(left, right)
}

func main() {
    arr := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
    ch := make(chan []int)

    go concurrentMergeSort(arr, ch)

    sorted := <-ch

    fmt.Println("Sorted array:", sorted)
}
