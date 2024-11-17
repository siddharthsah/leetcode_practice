package main

import "fmt"

// Time Complexity: O(log N)
// The binary search algorithm halves the search space in each iteration, leading to a logarithmic time complexity.   

// Space Complexity: O(1)
// The algorithm uses a constant amount of extra space, regardless of the input array size. 

func binarySearchTarget(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := (right + left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid -1
		}
	}
	return -1
}

func main() {
	nums := []int{ -1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(binarySearchTarget(nums, target))
}