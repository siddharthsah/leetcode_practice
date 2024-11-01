package main

import (
	"fmt"
	"log"
	"time"
)

var nums = []int{2, 7, 10, 3}
var target int = 12

func twoSum(nums []int, target int) []int {

	var record = make(map[int]int)
	//result := make([]int,0,len(nums))
	var result []int

	for key, value := range nums {
		if _, ok := record[target-value]; ok {
			result = append(result, key)
			result = append(result, record[target-value])
			break
		}

		record[value] = key
	}

	return result
}

func main() {

	start := time.Now()

	fmt.Println("Answer is:", twoSum(nums, target))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
