package main

import (
    
    "fmt"
  
)

func main() {
    var n,m int
    fmt.Scan(&n,&m)
    cnt := n/m
    ans := n
    day := n+1
    for cnt>0{
        ans++
        if day%m==0{
            cnt++
        }
        day++
        cnt--
    }
    fmt.Println(ans)
    
}
