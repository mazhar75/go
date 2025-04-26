package main

import "fmt"
type node struct{
	value int
	next *node
}
var head *node
func insert(x int){

	if head == nil{
		head = &node{value : x,}
	} else {
		var cur *node = head
		for cur.next != nil{
            cur=cur.next
		}
		cur.next=&node{value:x,}
	}	
}
func print(){
    var cur *node=head
	for cur != nil{
		fmt.Print(cur.value)
		cur=cur.next
	}

}
func main(){
     var n int
	 fmt.Scan(&n)
	 a := make([]int,n)
	 for i :=0;i<n;i++{
		fmt.Scan(&a[i])
		insert(a[i])
	 }
	 fmt.Println(a)
	 print()

}