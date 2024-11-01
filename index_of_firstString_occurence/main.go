package main

import "fmt"
//this is a brute force which has time complexioty of O(nxm), with KMP, it will be O(n+m)
func index_of_first_occurence(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println("the index of first occurence is:", index_of_first_occurence(haystack, needle))
}