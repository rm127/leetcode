package main

func checkRecord(s string) bool {
	var prev byte
	lCount := 0
	aCount := 0
	for i := 0; i < len(s); i++ {
		l := s[i]
		if l == 'A' {
			aCount++
			if aCount > 1 {
				return false
			}
		}
		if l == 'L' && prev == 'L' {
			lCount++
			if lCount > 1 {
				return false
			}
		} else {
			lCount = 0
		}
		prev = l
	}
	return true
}
