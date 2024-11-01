package main

import "fmt"

type MyStack struct {
	queue []int 
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	tmp := this.queue[len(this.queue) - 1]
	this.queue = this.queue[: len(this.queue) - 1]
	return tmp
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.queue[len(this.queue) - 1]
}

func main() {
	Queue := Constructor()
	Queue.Push(3)
	Queue.Push(7)
	Queue.Push(4)
	fmt.Println(Queue.queue)
	fmt.Println("Top value", Queue.Top())
	fmt.Println("Pop value", Queue.Pop())
	fmt.Println(Queue.queue)
}