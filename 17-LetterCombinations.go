package main

func LetterCombinations(digits string) []string {
	dict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	var sets []string
	// total results
	size := 1

	for _, digit := range digits {
		chars := dict[string(digit)]
		size = size * len(chars)
		sets = append(sets, chars)
	}

	thresh := size
	// create the results array with empty values
	res := make([]string, size)

	// loops through all sets
	for _, set := range sets {
		// times to repeat the same letter. Get's divided by the letters in a set for every set
		thresh = thresh / len(set)
		// counter for repeated letters
		counter := 0
		// the letter in question
		index := 0

		// loops across the whole array
		for i := 0; i < size; i++ {
			res[i] = res[i] + string(set[index])
			counter++
			// change letter
			if counter >= thresh {
				index++
				counter = 0
			}
			// cycles the letters in a set
			if index >= len(set) {
				index = 0
			}
		}
	}

	return res
}
