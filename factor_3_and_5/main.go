package main

import "fmt"

func factor_3_5(n int)[]int {
	factors := []int{}
	if n <= 0 {
		return nil
	}
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			if i % 3 == 0 || i % 5 == 0 {
				factors = append(factors, i)
			}
		} 
	}
	return factors
}
//Time Complexity: O(n)
//Space Complexity: O(k), where k is the number of factors of 3 or 5.

func main(){
	n := 30
	fmt.Println(factor_3_5(n))
}