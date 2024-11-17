package main

import "fmt"
//in place modification makes in O(1) for space complexity
func reverseString(word []byte) []byte {
	left, right := 0, len(word)-1
	if len(word) <= 1 {
		return word
	}
	for left < right {
		word[left], word[right] = word[right], word[left]
		left++
		right--
	}
	return word
}

func main() {
	word := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(string(reverseString(word)))
}
