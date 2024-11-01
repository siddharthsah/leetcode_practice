package main

import (
	"fmt"
	"time"
)

func replaceGreatestElem(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	//find the max element from right hand side
	maxRight := nums[n-1]

	for i := n-2; i >= 0; i-- {
		temp := nums[i]
		nums[i] = maxRight
		if temp > maxRight {
			maxRight = temp
		}
		nums[n-1] = -1
	} 
	return nums
} 

func main() {
	nums := []int{17, 18, 5, 4, 6}
	start := time.Now()
	result := replaceGreatestElem(nums)
	elapsed := time.Since(start)
	fmt.Println("result is", result)
	fmt.Printf("time taken for the program to run is %s", elapsed)
}