package main

func RomanToInteger(s string) int {
	dict := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	res := 0

	for i := 0; i < len(s); i++ {
		curr := s[i]

		var next string
		if i+1 < len(s) {
			next = string(s[i+1])
		}

		match := false

		if next != "" {
			if curr == 'I' && next == "V" {
				res += 4
				match = true
			}
			if curr == 'I' && next == "X" {
				res += 9
				match = true
			}
			if curr == 'X' && next == "L" {
				res += 40
				match = true
			}
			if curr == 'X' && next == "C" {
				res += 90
				match = true
			}
			if curr == 'C' && next == "D" {
				res += 400
				match = true
			}
			if curr == 'C' && next == "M" {
				res += 900
				match = true
			}
		}
		if match {
			i++
		} else {
			res += dict[rune(curr)]
		}
	}

	return res
}
