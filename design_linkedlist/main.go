package main 

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}
func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *linkedList) Delete(key int) {

	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *node = nil
	curr := l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		fmt.Println("Key Not found")
		return
	}
	prev.next = curr.next
	l.length--
	fmt.Println("Node Deleted")
}

func main() {
	list := linkedList{}
	node1 := &node{data: 20}
	node2 := &node{data: 30}
	node3 := &node{data: 40}
	node4 := &node{data: 50}
	node5 := &node{data: 70}
	list.PushBack(node1)
	list.PushBack(node2)
	list.PushBack(node3)
	list.PushBack(node4)
	list.PushBack(node5)
	fmt.Println("Length = ", list.Len())
	list.Display()
	list.Delete(40)
	fmt.Println("Length = ", list.Len())
	list.Display()

}