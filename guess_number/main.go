package main

import "fmt"

// This binary search approach ensures that the search space is halved in each iteration, leading to a time complexity of O(log n).

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)
		if res > 0 {
			left = mid + 1
		} else if res < 0 {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1 //If the loop completes without finding the secret number, return -1 to indicate failure.
}

func main() {
	// Assuming the secret number is 6
	result := guessNumber(10)
	fmt.Println(result) // Output: 6
}
