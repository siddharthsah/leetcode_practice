package main

import (
	"fmt"
	"time"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	index := n-1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[index] = leftSquare
			left++
		} else {
			result[index] = rightSquare
			right--
		}
		index--
	}
	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	start := time.Now()
	squares := sortedSquares(nums)
	fmt.Println("Answer is:", squares)
	elapsed := time.Since(start)
	fmt.Printf("the program took %s", elapsed)
	
}