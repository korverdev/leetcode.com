package main

import (
	"testing"
)

func TestIntToRoman1(t *testing.T) {
	const expected = "MMMDCCXLIX"
	result := intToRoman(3749)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestIntToRoman2(t *testing.T) {
	const expected = "LVIII"
	result := intToRoman(58)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestIntToRoman3(t *testing.T) {
	const expected = "MCMXCIV"
	result := intToRoman(1994)

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(b.N)
	}
}
