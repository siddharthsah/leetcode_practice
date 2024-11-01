package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	//find the middle node
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//Reverse the second half
	secondHalf := reverse(slow)

	firstHalf := head
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next :=curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	// Create a sample linked list
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 2}
    head.Next.Next.Next = &ListNode{Val: 1}
	start := time.Now()
	if isPalindrome(head) {
		fmt.Println("the linked list is a palindrome.")
	} else {
		fmt.Println("its not a palindrome.")		
	}
	elapsed := time.Since(start)
	fmt.Printf("program took %s", elapsed)
}
