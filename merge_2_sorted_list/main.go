package main

import (
	"fmt"
	"time"
)

// ListNode is a representation of a node in a Linked list
type ListNode struct  {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for list1!= nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	} 
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummyHead.Next
}


func main() {
	start := time.Now()
	// Example usage
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergedList := mergeTwoLists(list1, list2)

	for mergedList != nil {
		fmt.Print(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
	fmt.Println()

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)

}
