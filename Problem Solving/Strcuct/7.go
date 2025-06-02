/*
 একটি প্রোগ্রাম লিখুন যেখানে Formatter নামক ইন্টারফেসে একটি Format() মেথড আছে যা কিছু স্ট্রিং এর একটি স্লাইস নেয় এবং একটি ফর্ম্যাট করা স্ট্রিং রিটার্ন করে। Formatter নামক ইন্টারফেসটিকে 
 CSVFormatter struct এবং HTMLFormatter struct দ্বারা ইমপ্লিমেন্ট করুন। ৷ Format নামে একটি ফাংশন তৈরি করুন যা একটি Formatter আর্গুমেন্ট এবং স্ট্রিংগুলির একটি স্লাইস নেয় এবং Format() 
 মেথডকে কল করে। একটি CSVFormatter এবং একটি HTMLFormatter দিয়ে ফাংশনটি পরীক্ষা করুন৷ আর্গুমেন্ট হিসেবে পাস করা Formatter টি একটি CSVFormatter নাকি HTMLFormatter তা নির্ধারণ করতে টাইপ 
 অ্যাসারশন ব্যবহার করুন এবং ঐটার নিজস্ব অতিরিক্ত মেথডকে কল করুন৷
 */

package main

import (
	"fmt"
	"strings"
)

// Formatter interface
type Formatter interface {
	Format([]string) string
}

// CSVFormatter struct
type CSVFormatter struct{}

func (c CSVFormatter) Format(items []string) string {
	return strings.Join(items, ",")
}

func (c CSVFormatter) ExtraInfo() string {
	return "CSVFormatter is formatting as comma-separated values."
}

// HTMLFormatter struct
type HTMLFormatter struct{}

func (h HTMLFormatter) Format(items []string) string {
	var builder strings.Builder
	builder.WriteString("<ul>\n")
	for _, item := range items {
		builder.WriteString(fmt.Sprintf("  <li>%s</li>\n", item))
	}
	builder.WriteString("</ul>")
	return builder.String()
}

func (h HTMLFormatter) ExtraInfo() string {
	return "HTMLFormatter is formatting as an unordered HTML list."
}

// Format function
func Format(f Formatter, items []string) {
	fmt.Println("Formatted Output:")
	fmt.Println(f.Format(items))

	// Type assertion
	switch formatter := f.(type) {
	case CSVFormatter:
		fmt.Println("Extra Info:", formatter.ExtraInfo())
	case HTMLFormatter:
		fmt.Println("Extra Info:", formatter.ExtraInfo())
	default:
		fmt.Println("Unknown Formatter Type")
	}
}

func main() {
	items := []string{"apple", "banana", "cherry"}

	csvFormatter := CSVFormatter{}
	htmlFormatter := HTMLFormatter{}

	fmt.Println("Using CSVFormatter:")
	Format(csvFormatter, items)
	fmt.Println()

	fmt.Println("Using HTMLFormatter:")
	Format(htmlFormatter, items)
}
