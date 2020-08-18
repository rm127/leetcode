package main

import "fmt"

func isNumber(s string) bool {
	dict := make(map[byte]int, 10)

	// space check
	seenNonSpace := false

	for i := 0; i < len(s); i++ {
		curr := s[i]

		// letter check
		if !isOperator(curr) && !isDigit(curr) {
			fmt.Println("-> letter check")
			return false
		}

		// space check
		if i > 0 && s[i-1] == ' ' && curr != ' ' && seenNonSpace && dict[' '] > 0 {
			fmt.Println("-> space check")
			return false
		}
		if !seenNonSpace && curr != ' ' {
			seenNonSpace = true
		}

		// exp required digits check
		if curr == 'e' {
			if i+1 < len(s) && i > 0 {
				if !isDigit(s[i-1]) || !(isDigit(s[i+1]) || (isOperator(s[i+1]) && s[i+1] != ' ')) {
					fmt.Println("-> exp check")
					return false
				}
			} else {
				fmt.Println("-> exp check")
				return false
			}
		}

		if isOperator(curr) {
			if curr != ' ' && dict[curr] > 0 {
				fmt.Println("-> same twice")
				return false
			}
			dict[curr]++
		}
	}
	return true
}

func isDigit(s byte) bool {
	return s == '0' ||
		s == '1' ||
		s == '2' ||
		s == '3' ||
		s == '4' ||
		s == '5' ||
		s == '6' ||
		s == '7' ||
		s == '8' ||
		s == '9'
}

func isOperator(s byte) bool {
	return s == '+' ||
		s == '-' ||
		s == 'e' ||
		s == '.' ||
		s == ' '
}
