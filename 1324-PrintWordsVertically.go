package main

import "fmt"

func printVertically(s string) []string {
	if s == "" {
		return []string{}
	}

	// get size of longest word to set number of rows in slice
	size := 0
	count := 0
	maxWords := 0
	for i := 0; i < len(s)+1; i++ {
		if i < len(s) && s[i] == ' ' || i == len(s) {
			if count > size {
				size = count
				maxWords = 0
			}
			if count == size {
				maxWords++
			}
			count = 0
		} else {
			count++
		}
	}

	// create said array
	res := make([]string, size)

	rowIndex := 0
	curWordSize := 0
	passedMaxWords := 0
	for i := 0; i < len(s); i++ {
		letter := s[i]

		// if the letter is a space, the current word is done
		if letter == ' ' {
			// if the current word is shorter than the longest
			// and there are more maxLength words
			// if no more maxLength words are left, skip adding (trailing) spaces
			fmt.Println("space here", curWordSize, size, passedMaxWords, maxWords)
			if curWordSize < size && passedMaxWords < maxWords {
				// fill diff with spaces
				for rowIndex < size {
					res[rowIndex] += " "
					rowIndex++
				}
			} else {
				passedMaxWords++
			}
			// reset row and word size
			curWordSize = 0
			rowIndex = 0
		} else {
			curWordSize++
			// add current letter to res
			res[rowIndex] += string(letter)
			// loop incr
			if rowIndex == size-1 {
				rowIndex = 0
			} else {
				rowIndex++
			}
		}
	}
	return res
}
