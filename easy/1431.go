package main

import (
	"sort"
)

/* Runtime: 0ms 100%
* Memory: 2.4MB ~13%
**/
func kidsWithCandiesSolution1(candies []int, extraCandies int) []bool {
	l := len(candies)

	candiesSorted := make([]int, l)

	copy(candiesSorted, candies)

	var extraSlice []bool

	sort.Ints(candiesSorted)

	greatest := candiesSorted[l-1]

	for _, value := range candies {
		extraSlice = append(extraSlice, (value+extraCandies) >= greatest)
	}

	return extraSlice
}

/* Runtime: 0ms 100%
* Memory: 2.3MB ~13%
**/
func kidsWithCandiesSolution2(candies []int, extraCandies int) []bool {
	l := len(candies)

	var extraSlice []bool

	greatest := candies[0]

	for i := 1; i < l; i++ {
		if candies[i] > greatest {
			greatest = candies[i]
		}
	}

	for i := 0; i < l; i++ {
		extraSlice = append(extraSlice, (candies[i]+extraCandies) >= greatest)
	}

	return extraSlice
}

func main() {
	testCaseCandies := []int{2, 3, 5, 1, 3}
	testCaseExtra := 3

	kidsWithCandiesSolution1(testCaseCandies, testCaseExtra)
}
