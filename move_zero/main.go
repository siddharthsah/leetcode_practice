package main

import (
	"fmt"
	"time"
)

func moveZero(nums []int) {
	n := len(nums)

	slow := 0

	for fast := 0; fast < n; fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
// Time Complexity: O(n)
// The loop iterates through the array once, making it a linear time complexity.

// Space Complexity: O(1)
// The algorithm modifies the input array in-place, without using any extra space proportional to the input size.

func main() {
	start := time.Now()
	nums := []int{0, 1, 0, 3, 12}
	moveZero(nums)
	fmt.Println("after moving zeros to the right:", nums)
	elapsed := time.Since(start)
	fmt.Printf("time the program took: %s", elapsed)

}