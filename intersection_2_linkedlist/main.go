package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// By adjusting the pointers in this way, we ensure that they will eventually meet at the intersection point or both reach the end of their respective lists if there's no intersection.
// This approach has a time complexity of O(M + N), where M and N are the lengths of the two lists, and a space complexity of O(1).

func intersection2LinkedList(headA *ListNode, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

func main() {
	// Create a sample linked list
    headA := &ListNode{Val: 1}
    headA.Next = &ListNode{Val: 2}
    headA.Next.Next = &ListNode{Val: 3}
    headA.Next.Next.Next = &ListNode{Val: 4}
	headA.Next.Next.Next.Next = &ListNode{Val: 5}
	headB := &ListNode{Val: 6}
    headB.Next = &ListNode{Val: 7}
    headB.Next.Next = &ListNode{Val: 8}
    headB.Next.Next.Next = headA.Next.Next
	headB.Next.Next.Next.Next = &ListNode{Val: 9}
	headB.Next.Next.Next.Next.Next = &ListNode{Val: 10}
	fmt.Println(intersection2LinkedList(headA, headB))


}