package main

import "fmt"

// this is brute force and will be o(n) for time complexity

// func arrangingCoins(n int) int {
// 	sumOfRows := 0
// 	for i := 1; i <= n; i++ {
// 		sumOfRows += i
// 		if sumOfRows < n {
// 			continue
// 		} else if sumOfRows > n {
// 			return i - 1
// 		} else {
// 			return i
// 		}
// 	}
// 	return sumOfRows
// }

// Time Complexity: O(log N)
// The binary search algorithm halves the search space in each iteration, leading to a logarithmic time complexity.

// Space Complexity: O(1)
// The algorithm uses a constant amount of extra space, regardless of the input size.
func arrangingCoins(n int) int {
	left, right := 0 , n
	for left <= right {
		mid := left + (right - left) / 2
		currSum := mid * (mid + 1) / 2

		if currSum == n {
			return mid
		} else if currSum > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	n := 5
	fmt.Println(arrangingCoins(n))
}