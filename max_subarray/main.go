package main

import (
	"fmt"
	"time"
)

func maxSubArray(nums []int) int {
	maxSoFar, maxEndingHere := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}




func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	start := time.Now()
	fmt.Println(maxSubArray(nums))
	elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
}
