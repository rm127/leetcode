package main

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

	rowIndex := size - 1
	curWordSize := 0
	// passedMaxWords := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]

		// if the letter is a space, the current word is done
		if letter == ' ' {
			if curWordSize < size {
				for rowIndex > 0 {
					res[rowIndex] = " " + res[rowIndex]
					rowIndex--
				}
			}
			curWordSize = 0
		} else {
			curWordSize++
			res[rowIndex] = string(letter) + res[rowIndex]
			if rowIndex == 0 {
				rowIndex = size - 1
			} else {
				rowIndex--
			}
		}
	}
	return res
}
