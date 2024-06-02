package main

import (
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	const expected = true
	result := isPalindrome(121)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestIsPalindrome2(t *testing.T) {
	const expected = false
	result := isPalindrome(-121)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestIsPalindrome3(t *testing.T) {
	const expected = false
	result := isPalindrome(10)

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	midpoint := b.N / 2

	for i := -midpoint; i < midpoint; i++ {
		isPalindrome(b.N)
	}
}
