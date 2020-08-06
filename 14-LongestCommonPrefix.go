package main

func LongestCommonPrefix(strs []string) string {
	lcp := ""
	incr := 0
	current := ""
	done := false

	if len(strs) == 0 {
		return ""
	}

	for !done {
		curLength := 0
		for i, set := range strs {
			curLength = len(set)
			if curLength == 0 || incr+1 > curLength {
				done = true
				break
			}
			// add a new letter to the prefix
			if i == 0 {
				current += string(set[incr])
			}
			// if mismatch, then signal done
			if current[incr] != set[incr] {
				done = true
			}
		}
		// if all matches, then increase and save the current prefix
		if !done {
			incr++
			lcp = current
		}
		// stop if we've reached the end of a set
		if incr == curLength {
			break
		}
	}

	return lcp
}
