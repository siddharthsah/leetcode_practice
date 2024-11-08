package main

import "fmt"

// Time Complexity: O(N)
// Space Complexity: O(1) (excluding the output array)

// This approach efficiently finds the missing numbers by utilizing the array itself to 
// mark visited elements, avoiding the need for additional data structures.

func findDisappearedNumber(nums []int) []int {
	len := len(nums)

	for i := 0; i < len; i++ {
		index := abs(nums[i]) - 1
		nums[index] = -abs(nums[index])
	}

	var missingArray []int
	for i := 0; i < len; i++ {
		if nums[i] > 0 {
			missingArray = append(missingArray, i+1)
		}
	}
	return missingArray
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	missing := findDisappearedNumber(nums)
	fmt.Println(missing)
}
