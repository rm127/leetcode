package main

func ValidParentheses(s string) bool {
	dict := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var path []rune
	for i, char := range s {
		size := len(s)
		if i < size {
			if len(path) > 0 && char == path[len(path)-1] {
				path = path[:len(path)-1]
			} else {
				path = append(path, dict[char])
			}
		}
	}
	return len(path) == 0
}
