package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool) // Use a map to detect cycles

	for n != 1 && !seen[n] {
			seen[n] = true
			sum := 0
			for n > 0 {
					digit := n % 10
					sum += digit * digit
					n /= 10
			}
			n = sum
	}

	return n == 1
}

func main() {
	n := 19
	if isHappy(n) {
		fmt.Println(n, "is a happy number")
	} else {
		fmt.Println(n, "is not a happy number")
	}
}

/*
Time and Space Complexity:

Time Complexity:
In the worst case (a non-happy number entering a cycle), the algorithm might iterate a number of times proportional to the input number n. However, 
the numbers quickly decrease due to the squaring of digits. Therefore, the time complexity is difficult to express with a tight bound using Big O 
notation in terms of the initial input n. In practice, it's observed that the number of iterations is relatively small even for large numbers. The 
number of possible cycles is also limited. If we consider the number of iterations to be bounded by a constant, the complexity can be considered O(1) 
in practice for the range of integers typically encountered. However, theoretically, it's not strictly constant.

Space Complexity:
The space complexity is determined by the seen map. In the worst case, the map might store a number of entries proportional to the input number n 
before a cycle is detected. However, as with time complexity, the numbers decrease rapidly, limiting the number of entries in the seen map.
Therefore, the space complexity can also be considered O(1) in practice. Again, the theoretical worst-case is not strictly constant, but it's 
bounded in practice.
*/