package main

import (
	"fmt"
	"strings"
)

/* _______ | Amount | Percentage |
 * Runtime | 0ms    | 100%     |
 * Memory  | 1.9MB  | 24.67%     |
**/
func defangIPaddrOne(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

/* _______ | Amount | Percentage |
 * Runtime | 4ms    | 94.65%     |
 * Memory  | 3.8MB  | 34.37%     |
**/
func defangIPaddrTwo(address string) string {
	str := ""

	for _, v := range address {
		if v == rune('.') {
			str += "[.]"
			continue
		}
		str += string(v)
	}
	return str
}
func main() {
	ip := "1.1.1.1"

	fmt.Println(defangIPaddrOne(ip))
	fmt.Println(defangIPaddrTwo(ip))
}
