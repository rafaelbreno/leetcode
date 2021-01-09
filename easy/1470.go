package main

import "fmt"

/* _______ | Amount | Percentage |
 * Runtime | 4ms    | 94.65%     |
 * Memory  | 3.8MB  | 34.37%     |
**/
func shuffleOne(nums []int, n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, nums[i])
		arr = append(arr, nums[i+n])
	}

	return arr
}

func shuffleTwo(nums []int, n int) []int {
	arr := make([]int, len(nums))
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		arr[j], arr[j+1] = nums[i], nums[i+n]
	}

	return arr
}

func main() {
	testNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testN := 5

	fmt.Println(testNums)
	fmt.Println(shuffleOne(testNums, testN))
	fmt.Println(shuffleTwo(testNums, testN))
}
