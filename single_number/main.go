package main

import "fmt"

// Time Complexity: O(N)
// The code iterates through the array once, making it a linear time complexity.

// Space Complexity: O(1)
// The code uses constant extra space, independent of the input array size.

// Why XOR Works:

// XORing a number with itself results in 0.
// XORing a number with 0 results in the number itself.
// Therefore, when we XOR all numbers in the array, the duplicate numbers will cancel each other out, leaving only the single number. 
// This is a concise and efficient approach to finding the single number.

func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func main() {
	nums := []int {4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}