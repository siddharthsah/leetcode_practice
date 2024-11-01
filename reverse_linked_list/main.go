package main

import (
	"fmt"
	"time"
)

// ListNode is a representation of a node in a Linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse_linked_list(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		temp_next := current.Next
		current.Next = prev
		prev = current
		current = temp_next
	}
	return prev
}

func main() {
	start := time.Now()
	// Example usage
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	reveresedList := reverse_linked_list(head)

	for reveresedList != nil {
		fmt.Print(reveresedList.Val, " ")
		reveresedList = reveresedList.Next
	}
	fmt.Println()

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s", elapsed)
}
