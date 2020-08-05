package main

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	answer := 0
	dict := make(map[int]int)
	i := 0
	for j := 0; j < n; j++ {
		charval := int(s[j])
		if val, ok := dict[charval]; ok && val != i {
			if val > i {
				i = val
			}
		}
		curval := j - i + 1
		if curval > answer {
			answer = curval
		}
		dict[charval] = j + 1
	}
	return answer
}
