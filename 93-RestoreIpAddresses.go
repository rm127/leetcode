package main

import (
	"math"
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	return RestoreIp(3, "", s)
}

func RestoreIp(rd int, p string, s string) []string {
	l := len(s)

	if len(p) > 1 && p[0] == '0' {
		// if postfix' first digit is zero with a non-zero digit after
		return []string{}

	} else if rd == 3 && l < 4 {
		// an Ip address requires at least 4 digits
		return []string{}

	} else if l > 3 && rd == 0 {
		// if there are no dots left and string is too long
		return []string{}

	} else if l <= 3 && l > 0 && rd == 0 {
		// if the element is 0 < x < 3 and no dots are left

		// if the element's first digit is zero with a non-zero digit after
		if len(s) > 1 && s[0] == '0' {
			return []string{}
		}

		// convert to number
		num, _ := strconv.Atoi(s)
		// if first number is above 255 limit
		if num > 255 {
			return []string{}
		}

		return []string{s + "." + p}
	}

	var res []string
	var options []string

	// get the last biggest number (3-1 digits)
	max := int(math.Min(float64(l), 3))

	// convert to number
	num, _ := strconv.Atoi(s[l-max:])

	for i := max; i > 0; i-- {
		// if number is above 255 limit
		if i == 3 && num > 255 {
			continue
		}
		options = append(options, RestoreIp(rd-1, s[l-i:], s[0:l-i])...)
	}

	// append all possible numbers
	for i := 0; i < len(options); i++ {
		value := options[i]
		// only if postfix is present
		if p != "" {
			value += "." + p
		}
		res = append(res, value)
	}

	return res
}
