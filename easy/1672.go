package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	richest := 0
	currentRichest := 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			currentRichest += accounts[i][j]
		}
		if currentRichest > richest {
			richest = currentRichest
			currentRichest = 0
		}
	}

	return richest
}

func main() {
	testCase := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}

	fmt.Println(maximumWealth(testCase))
}
