package main

import (
	"fmt"
	"time"
)

func is_anagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	} 

	count := [26]int{} //this is a hash maps which maintains the frequency of each character
	// count := make(map[rune]int) to handle Unicode characters in Go.  str := "Hello, 世界!"
	for _, c := range s {
		count[c- 'a']++
	}
	for _, c := range t {
		count[c- 'a']--
		if count[c- 'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "anagram"
    t1 := "nagaram"
    s2 := "rat"
    t2 := "car"
	start := time.Now()
	fmt.Println(is_anagram(s1, t1))
	
	fmt.Println(is_anagram(s2, t2))
	elapsed := time.Since(start)
	fmt.Printf("program took %s\n", elapsed)
}
