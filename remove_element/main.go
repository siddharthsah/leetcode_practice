package main

import "fmt"

func removeElement(nums []int, target int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			nums[k] = nums[i]
			k++
		}	
	}
	return k
}



func main() {
	nums := []int {1, 2, 4, 4, 5, 9, 5, 1, 4}
	target := 4

	fmt.Println("final lenght of the array is ", removeElement(nums, target))
}