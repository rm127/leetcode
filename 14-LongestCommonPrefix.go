package main

func LongestCommonPrefix(strs []string) string {
	lcp := ""
	incr := 0
	done := false

	if len(strs) == 0 {
		return ""
	}

	for !done {
		for _, set := range strs {
			curLength := len(set)
			// if current set isn't empty and isn't out of bounds
			if curLength == 0 || incr+1 > curLength {
				done = true
				break
			}
			// add a new letter to the prefix
			// if mismatch, then signal done
			if strs[0][incr] != set[incr] {
				done = true
			}
		}

		// if all matches
		if !done {
			// take the last letter from the first set and append to result
			lcp += string(strs[0][incr])
			incr++
		}
	}

	return lcp
}
