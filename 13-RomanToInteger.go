package main

var dict = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func RomanToInteger(s string) int {

	if len(s) == 0 {
		return 0
	}

	length := len(s)
	res := dict[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		curr := s[i]
		next := s[i+1]
		tempRes := dict[curr]

		if curr == 'I' && (next == 'V' || next == 'X') {
			tempRes *= -1
		} else if curr == 'X' && (next == 'L' || next == 'C') {
			tempRes *= -1
		} else if curr == 'C' && (next == 'D' || next == 'M') {
			tempRes *= -1
		}
		res += tempRes
	}

	return res
}
