package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	goods := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] == nums[j] {
				goods++
			}
		}
	}

	return goods
}

func main() {
	testCase1 := []int{1, 2, 3, 1, 1, 3}

	fmt.Println(numIdenticalPairs(testCase1))
}
