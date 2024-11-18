package main

import "fmt"

// Time Complexity:

// Iterating over the string: O(n), where n is the length of the text string.
// Calculating the minimum frequency: O(1), as it involves a constant number of comparisons.
// Therefore, the overall time complexity is O(n), dominated by the iteration over the string.

// Space Complexity:

// Character count map: O(1), as the number of unique characters in "balloon" is constant.
// Other variables: O(1) for the res variable.
// Therefore, the overall space complexity is O(1), indicating constant space usage.

func maxNumberOfBalloons(text string) int {
    count := map[byte]int{}
    for _, c := range text {
        count[byte(c)]++
    }

    return min(count['b'], count['a'], count['l'] / 2, count['o']/2, count['n'])
}


// func maxNumberOfBalloons(text string) int {
//     countText := counter(text)
//     balloon := counter("balloon")
    
//     res := len(text)
//     for c, _ := range balloon {
//         res = min(res, countText[c] / balloon[c])
//     }
//     return res
// }

// func counter(text string) map[rune]int {
//     count := make(map[rune]int)
//     for _, c := range text {
//         count[c] = count[c] + 1
//     }
//     return count
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

func main() {
	text := "loonbalxballpoon"
	fmt.Println(maxNumberOfBalloons(text))
}