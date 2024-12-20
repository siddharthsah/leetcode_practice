package main

import (
	"fmt"
	"strings"
)
// time complexity and space complexity is O(n+m)
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	seenChr, seenWord := make(map[byte]int, len(pattern)), make(map[string]int, len(pattern))
	for i := 0; i < len(pattern); i++ {
		chrPos, chrOk := seenChr[pattern[i]]
		wordPos, wordOk := seenWord[words[i]]
		if chrOk != wordOk || chrPos != wordPos {
			return false
		}

		seenChr[pattern[i]] = i
		seenWord[words[i]] = i
	}

	return true
}

func main() {
	pattern := "abba"
	s := "dog bat bat dog"
	if wordPattern(pattern, s) {
		fmt.Println("The strings follow the pattern")
	} else {
		fmt.Println("The strings do not follow the pattern")
	}

}
