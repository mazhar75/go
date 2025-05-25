/*
একটি URL এর লিস্ট থাকবে। গো-রুটিন ব্যবহার করে প্রত্যেকটি URL এর ডাটা ডাউনলোড করতে হবে এবং একটি Struct অ্যারেতে ডাটাগুলো স্টোর করে রাখতে হবে। 
চ্যানেল ব্যবহার করে কাজটি ধারাবাহিক ভাবে করতে হবে।
*/

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type DownloadedData struct {
    URL  string
    Data []byte
}

func downloadURL(url string, ch chan DownloadedData) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error downloading %s: %v\n", url, err)
        ch <- DownloadedData{URL: url, Data: nil}
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading body %s: %v\n", url, err)
        ch <- DownloadedData{URL: url, Data: nil}
        return
    }

    ch <- DownloadedData{URL: url, Data: body}
}

func main() {
    urls := []string{
        "https://example.com",
        "https://golang.org",
        "https://google.com",
    }

    ch := make(chan DownloadedData)

    for _, url := range urls {
        go downloadURL(url, ch)
    }

  
    var results []DownloadedData
    for range urls {
        data := <-ch
        results = append(results, data)
    }

    for _, d := range results {
        fmt.Printf("URL: %s, Data size: %d bytes\n", d.URL, len(d.Data))
    }
}
