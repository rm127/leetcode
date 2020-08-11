package main

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	seenSegment := false
	seenGap := false
	for i, char := range s {
		// check if first character is part of a segment = increase count
		if i == 0 && char != ' ' {
			count++
		}
		if char == ' ' {
			seenGap = true
		} else {
			// if we've seen a segment before and a gap before and now see a segment again = increase count
			if seenSegment && seenGap {
				seenSegment = false
				seenGap = false
				count++
			}
			seenSegment = true
		}
	}
	return count
}
