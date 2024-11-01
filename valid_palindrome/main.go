package main

import (
	"fmt"
	"strings"
)



// // 8. Strings
// var s string = "this is a string"
// // treated like an array (collection or array)
// fmt.Printf("Value = %v, Type = %T\n", s, s)
// fmt.Printf("Value = %v, Type = %T\n", s[0], s[0])
// // s[2] = 'o' //cannot assign as string are immutable in go
// // concat using +

// // convert them into a collection/slice of bytes
// b := []byte(s)
// fmt.Printf("Value = %v, Type = %T\n", b, b)

// // 9. Rune (utf-32 characters)
// // Type alias for int32
// var r rune = 'a'
// fmt.Printf("Value = %v, Type = %T\n", r, r)


func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1

	for left <right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]){
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println("the string is palindrome:",isPalindrome(s))
}