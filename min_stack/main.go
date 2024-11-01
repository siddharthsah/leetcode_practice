package main

import "fmt"

type MinStack struct {
	data []int
	min []int
}

func Constructor() *MinStack {
	return &MinStack{
		data: make([]int, 0),
		min: make([]int, 0),
	}
}

func (stack *MinStack) Push(val int) {
	stack.data = append(stack.data, val)
	if len(stack.min) == 0 || val <= stack.min[len(stack.min)-1] {
		stack.min = append(stack.min, val)
	}
}

func (stack *MinStack) Pop() {
	if len(stack.data) > 0 {
		if stack.data[len(stack.data)-1] == stack.min[len(stack.min)-1] {
			stack.min = stack.min[:len(stack.min)-1]
		}
		stack.data = stack.data[:len(stack.data)-1]
	}
}

func (stack *MinStack) Top() int {
	if len(stack.data) > 0 {
		return stack.data[len(stack.data)-1]
	}
	return -1
}

func (stack *MinStack) GetMin() int {
	if len(stack.min) > 0 {
		return stack.min[len(stack.min)-1]
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.GetMin()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.GetMin()
	obj.Pop()
	obj.Top()
	obj.GetMin()
	fmt.Println(obj)
}