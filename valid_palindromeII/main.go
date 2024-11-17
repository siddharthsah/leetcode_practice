package main

import "fmt"

func isPalindromeII(word string) bool {
	left, right := 0, len(word) - 1
	for left < right {
		if word[left] != word[right] {
			return isPalindrome(word[left + 1 : right]) || isPalindrome(word[left : right -1]) 
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	} 
	return true
}

func main() {
	word := "abca"
	fmt.Println(isPalindromeII(word))
}