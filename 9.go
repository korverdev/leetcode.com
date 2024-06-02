// https://leetcode.com/problems/palindrome-number/
package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	// Convert the int to a string
	x_str := strconv.Itoa(x)
	length_of_x := len(x_str)

	halfway := length_of_x / 2

	for i := 0; i < halfway; i++ {
		if x_str[i] != x_str[length_of_x-i-1] {
			return false
		}
	}

	return true
}
