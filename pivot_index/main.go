package main

import "fmt"
// Time Complexity: O(N)
// The code iterates through the array once, making it a linear time complexity.

// Space Complexity: O(1)
// The code uses a constant amount of extra space, regardless of the input array size.

// This approach efficiently finds the pivot index by calculating the left sum and comparing it to the right sum in a single pass through the array.
func findPivotIndex(nums []int) int {
	totalSum := 0

	leftSum := 0

	for _, num := range nums {
		totalSum += num
	}
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

func main(){
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(findPivotIndex(nums))
}