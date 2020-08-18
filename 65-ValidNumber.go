package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	fmt.Println(strings.Join(parseNumber(s), ", "))
	return true
}

func parseNumber(s string) []string {
	if len(s) > 0 {
		curr := s[0]
		if !isDigit(curr) && !isOperator(curr) {
			return append(parseNumber(s[1:]), "letter")
		}
		if curr == ' ' {
			return append(parseNumber(s[1:]), "space")
		}
		if curr == 'e' {
			return append(parseNumber(s[1:]), "exp")
		}
		for i := 1; i < len(s); i++ {
			if (!isDigit(curr) && !isOperator(curr)) || curr == ' ' || curr == 'e' || i == len(s)-1 {
				fmt.Println(i, string(curr), "-"+s[0:i]+"-")
				return append(parseNumber(s[i+1:]), "xx")
			}
		}
	}
	return []string{}
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
