package main

import "fmt"

// this code has O( n x m) where n and m is the length of nums 1 and nums 2 

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	nums1Index := make(map[int]int)
// 	for index, val := range nums1 {
// 		nums1Index[val] = index
// 	}
// 	result := []int{}
// 	for i, v := range nums1 {
// 		nums1Index[v] = i
// 		result = append(result, -1)
// 	}

// 	for i := range len(nums2) {
// 		if _, ok := nums1Index[nums2[i]]; !ok {
// 			continue
// 		}

// 		for j := range nums2[i+1:] {
// 			if nums2[j] > nums2[i] {
// 				index := nums1Index[nums2[i]]
// 				result[index] = nums2[j]
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

// Time Complexity: O(N + M), where N is the length of nums1 and M is the length of nums2.
// Space Complexity: O(M), for the m map to store the indices of elements in nums2.

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1idx := make(map[int]int)
	ans := []int{}
	for i, v := range nums1 {
		nums1idx[v] = i
		ans = append(ans, -1)
	}
	stack := []int{}

 	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			idx := nums1idx[val]
			ans[idx] = v
		}

		if _, e := nums1idx[v]; !e {
			continue
		}
		stack = append(stack, v)
	}
	return ans
}
func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
