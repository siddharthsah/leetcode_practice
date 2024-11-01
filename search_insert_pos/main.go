package main

import (
	"fmt"
	"time"
)

func searchInsertPosition(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	start := time.Now()
	index := searchInsertPosition(nums, target)
	elapsed := time.Since(start)
	fmt.Printf("program took %s", elapsed)
	fmt.Println("Insert position:", index)
}
