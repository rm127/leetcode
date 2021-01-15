package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits) - 1
	// if the last digit is smaller than 9, just add 1
	if digits[length] < 9 {
		digits[length]++
		return digits
	}

	// if input is [9], then return [1,0]
	if length == 0 && digits[length] == 9 {
		return []int{1, 0}
	}

	// increase next number if current == 9. Stop at second first entry
	for digits[length] >= 9 && length > 0 {
		digits[length] = 0
		if length > 0 {
			digits[length-1]++
		}
		length--
	}
	fmt.Println(digits)
	// check for first number == 10
	if digits[0] == 10 {
		digits[0] = 0
		return append([]int{1}, digits...)
	}
	return digits
}
