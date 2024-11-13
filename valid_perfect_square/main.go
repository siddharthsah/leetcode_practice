package main

import "fmt"

// Time Complexity: O(log N)
// Space Complexity: O(1)


func validPerfectSquare(n int) bool {
	left, right := 1, n
	for left <= right {
		mid := left + (right - left) / 2
		if mid * mid > n {
			right = mid - 1
		} else if mid * mid == n {
			return true
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	n := 14
	fmt.Println(validPerfectSquare(n))
}