package main

import (
	"fmt"
	"time"
)

func house_robber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	rob1, rob2 := nums[0], max(nums[0],nums[1]) // using 2 variable instead of whole array for dp where 2 variable are initializing the robs from adjacent houses
	for i := 2; i < n; i++ {
		curr := max(rob2, rob1+nums[i])
		rob1, rob2 = rob2, curr
	}
	return rob2
}

func main() {
	nums := []int{1, 2, 3, 1,3}
	start := time.Now()
	fmt.Println(house_robber(nums))
	elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
}