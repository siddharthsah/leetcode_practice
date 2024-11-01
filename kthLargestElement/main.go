package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}
//this function initiliazes the heap 
func Constructor(k int, nums []int) KthLargest {
	tmp := IntHeap(nums)
	number := KthLargest{minHeap: &tmp, k: k}
	heap.Init(number.minHeap)
	for len(*number.minHeap) > k {
		heap.Pop(number.minHeap)
	}
	return number
}
//returns kth largest element after adding the val
func (number *KthLargest) Add(val int) int {
	heap.Push(number.minHeap, val)
	if len(*number.minHeap) > number.k {
		heap.Pop(number.minHeap)
	}
	return (*number.minHeap)[0]

}

func main() {
	k := 3
	nums := []int{4, 5, 8, 2}
	KthLargest := Constructor(k, nums)
	fmt.Println(*KthLargest.minHeap)
	fmt.Println(KthLargest.Add(3)) // Output: 4
	fmt.Println(KthLargest.Add(1)) // Output: 4
	fmt.Println(KthLargest.Add(5))
}
