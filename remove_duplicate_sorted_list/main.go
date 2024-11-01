package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, curr := head, head.Next

	for curr != nil {
		for curr != nil && curr.Val == prev.Val {
			curr = curr.Next
			prev.Next = curr
		}

		if curr != nil {
			prev, curr = curr, curr.Next
		}
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
			fmt.Printf("%d -> ", current.Val)
			current = current.Next

	}
	fmt.Println("nil")
}

func main() {
	// Create a sample linked list
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	start := time.Now()
	final := removeDuplicate(head)
	PrintLinkedList(final)
	
	elapsed := time.Since(start)
	fmt.Printf("program took %s", elapsed)
}
