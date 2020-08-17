package main

func numUniqueEmails(emails []string) int {
	dict := make(map[string]int)

	for i := 0; i < len(emails); i++ {
		// increase the counter for the given email
		dict[sanitize(emails[i], false)] += 1
	}
	// return the count of items in the map
	return len(dict)
}

// removes dots and skips letters after + until @ is reached
func sanitize(s string, skipping bool) string {
	letter := s[0]
	// if @ just return the rest
	if letter == '@' {
		return s
	}
	// Skip the next letters until @
	if letter == '+' || skipping {
		return sanitize(s[1:], true)
	}
	// skip the current letter
	if letter == '.' {
		return sanitize(s[1:], false)
	}
	// return current letter and sanitize the rest
	return string(letter) + sanitize(s[1:], false)
}
