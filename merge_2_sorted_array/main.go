package main

import (
	"fmt"
	"time"
)

func mergeTwoArray(nums1 []int, m int, nums2 []int, n int) []int {
	p1, p2, last := m-1, n-1, m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[last] = nums1[p1]
			p1--
		} else {
			nums1[last] = nums2[p2]
			p2--
		}
		last--
	}

	for p2 >= 0 {
		nums1[last] = nums2[p2]
		p2--
		last--
	}
	return nums1
}

func main(){
	start := time.Now()
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6} 
	m := 3
	n := 3
	fmt.Println(mergeTwoArray(nums1, m, nums2, n))
	elapsed := time.Since(start)
	fmt.Printf("time elapsed since start of the program %s", elapsed)

}
