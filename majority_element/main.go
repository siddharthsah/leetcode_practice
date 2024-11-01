package main

import "fmt"

// majority element is the element that appears more than [n / 2] times. this is Boyern Moore which happens in o(1) space but if we use 
// hashify to do this, it will take o(n) space 

func majorityElement(nums []int) int {
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count += 1
		} else {
			count -= 1
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}
	return candidate
}

func main() {
	nums := []int {2, 2, 1, 1, 1, 2, 2}
	majority := majorityElement(nums)
	fmt.Println("Majority Element:", majority)
}