package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeLinkedList(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}
	return dummyHead.Next
}

func main() {
	// Create a sample linked list
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 6}
    head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	newHead := removeLinkedList(head, 6)

	for newHead != nil {
		fmt.Println(newHead.Val, " ")
		newHead = newHead.Next
	}
}