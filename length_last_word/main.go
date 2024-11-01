package main

import (
	"fmt"
	"time"
)

func lengthLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	wordFound := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
			wordFound = true
		} else if wordFound {
			break
		}
	}
	return count
}

func main() {
	s := " my name is jagrit  i "
	start := time.Now()
	fmt.Println("length of last word is ",lengthLastWord(s))
	elapsed := time.Since(start)
	fmt.Printf("time taken to run the program was %s", elapsed)
}
