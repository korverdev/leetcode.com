// https://leetcode.com/problems/integer-to-roman/
package main

import (
	"strings"
)

type symbol struct {
	value int
	roman string
}

func intToRoman(num int) string {
	symbols := []symbol{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	builder := strings.Builder{}

	for _, symbol := range symbols {
		for num >= symbol.value {
			num -= symbol.value
			builder.WriteString(symbol.roman)
		}
	}

	return builder.String()
}
