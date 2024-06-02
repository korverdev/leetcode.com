// https://leetcode.com/problems/palindrome-number/
package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	x_str := strconv.Itoa(x)
	length_of_x := len(x_str)

	halfway := length_of_x / 2
	end_index := length_of_x - 1

	for index := 0; index < halfway; index++ {
		if x_str[index] != x_str[end_index-index] {
			return false
		}
	}

	return true
}
