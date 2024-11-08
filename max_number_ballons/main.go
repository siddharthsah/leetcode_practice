package main

import "fmt"

// func counter(text string) map[rune]int {
// 	count := make(map[rune]int)
// 	for _, c := range text {
// 		count[c]++
// 	}
// 	return count
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func maxNumberBallons(text string) int {
// 	countText := counter(text)
// 	ballon := counter("ballon")

// 	res := len(text)
// 	for _, c := range text {
// 		res = min(res, countText[c] / ballon[c])
// 	}
// 	return res
// }
// func main() {
// 	text := "loonbalxballpoon"
// 	fmt.Println(maxNumberBallons(text))
// }

func maxNumberOfBalloons(text string) int {
    countText := counter(text)
    balloon := counter("balloon")
    
    res := len(text)
    for c, _ := range balloon {
        res = min(res, countText[c] / balloon[c])
    }
    return res
}

func counter(text string) map[rune]int {
    count := make(map[rune]int)
    for _, c := range text {
        count[c] = count[c] + 1
    }
    return count
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}