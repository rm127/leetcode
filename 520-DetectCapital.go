package main

import (
	"unicode"
)

func DetectCapital(word string) bool {
	if len(word) == 1 {
		return true
	}
	return DetectCap(word, false, false, 0)
}

func DetectCap(word string, firstIsUpper bool, secondIsUpper bool, count int) bool {
	isUpper := unicode.IsUpper(rune(word[0]))

	if count > 0 && !firstIsUpper && isUpper {
		// if 1st letter is lower and 2nd is upper -> after the first 2 letters
		return false
	} else if count > 1 && firstIsUpper && !secondIsUpper && isUpper {
		// if 1st is upper, 2nd is lower and current is upper -> after the first 2 letters
		return false
	} else if firstIsUpper && secondIsUpper && !isUpper {
		// if 1st is upper, 2nd is upper and current is lower
		return false
	}

	if len(word) > 1 {
		if count == 0 {
			firstIsUpper = isUpper
		} else if count == 1 {
			secondIsUpper = isUpper
		}
		return DetectCap(word[1:], firstIsUpper, secondIsUpper, count+1)
	}
	return true
}
