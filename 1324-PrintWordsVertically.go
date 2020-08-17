package main

import (
	"strings"
)

func printVertically(s string) []string {
	if s == "" {
		return []string{}
	}

	words := strings.Split(s, " ")
	size := 0

	// get size of longest word to set number of rows in slice
	for i := 0; i < len(words); i++ {
		if len(words[i]) > size {
			size = len(words[i])
		}
	}

	// create said array
	res := make([]string, size)

	for i := 0; i < len(words); i++ {
		for j := 0; j < size; j++ {
			if j < len(words[i]) {
				res[j] += string(words[i][j])
			} else {
				res[j] += " "
			}

		}
	}

	// remove trailing whitespace
	for i := 0; i < len(res); i++ {
		for j := len(res[i]) - 1; j >= 0; j-- {
			item := res[i]
			if item[j] != ' ' {
				res[i] = res[i][0 : j+1]
				break
			}
		}
	}

	return res
}
