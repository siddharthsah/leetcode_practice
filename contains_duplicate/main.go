package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)  //create a hashmap which makes the memory complexity o(n) but decreaes the space from o(n2)(when checking using brute force) to o(n)
	// we could reduce the o(n2) to o(nlogn) if we sort the array but with hash map we can reduce to it o(n)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
} 

func main() {
	nums := []int {1, 4, 3, 6, 8, 9, 5, 2, 7}
	fmt.Println("the above array of integer contains duplicate:", containsDuplicate(nums))

}