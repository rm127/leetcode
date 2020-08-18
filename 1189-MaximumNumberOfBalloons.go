package main

import (
	"strings"
)

func maxNumberOfBalloons(text string) int {
	if text == "" {
		return 0
	}

	storage := make(map[byte]int, 5)

	s := strings.ToLower(text)

	for i := 0; i < len(s); i++ {
		curr := s[i]
		if curr == 'l' || curr == 'o' {
			storage[curr]++
		}
		// count the other letters twice, to match the double "ll" and "oo"
		if curr == 'b' || curr == 'a' || curr == 'n' {
			storage[curr] += 2
		}
	}

	if len(storage) < 5 {
		return 0
	}

	// set res way too high
	res := 99999

	// count how many times the letters occurs. Keep the smallest number
	for _, elem := range storage {
		if elem < res {
			res = elem
		}
	}
	// split by two due to the "ll" and "oo" fix.
	return res / 2
}
