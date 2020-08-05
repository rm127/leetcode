package main

func LongestPalindrome(s string) string {
	size := len(s)
	if size <= 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < size; i++ {
		even := longestFromCenter(s, i, i+1)
		odd := longestFromCenter(s, i, i)
		curlength := even
		if odd > even {
			curlength = odd
		}
		if (end - start) < curlength {
			start = i - (curlength - 1) / 2
			end = i + curlength / 2
		}
	}
	return s[start : end+1]
}

func longestFromCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}