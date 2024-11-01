package main

import (
	"fmt"
	"time"
)

func initList() *SingleList {
	return &SingleList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) AddFront(num int) {
	ele := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}
	s.Len++
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur := new(ListNode)
    ret := cur
    sum := 0
    for l1 != nil || l2 != nil {  
        if l1 != nil {
            sum = sum + l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum = sum + l2.Val
            l2 = l2.Next
        }
            
        cur.Next = &ListNode{
                sum%10,
                nil,
        }
        cur = cur.Next
        sum = sum/10
    }
    
    if sum > 0 {
        cur.Next = &ListNode{sum,nil}
    }
        
    return ret.Next
}

func main() {
	first := initList()	
    first.AddFront(2)
	first.AddFront(4)
	first.AddFront(3)

	second := initList()
	second.AddFront(5)
	second.AddFront(6)
	second.AddFront(4)
    start := time.Now()
	result := addTwoNumbers(first.Head, second.Head)
    elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
	result.Traverse()
}
