package main

import (
	"fmt"
	"math"
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	return RestoreIp(3, "", s)
}

func RestoreIp(rd int, p string, s string) []string {
	l := len(s)

	if rd == 3 && l < 4 {
		// an Ip address requires at least 4 digits
		return []string{}
	} else if l > 3 && rd == 0 {
		// if there are no dots left and string is too long
		return []string{}
	} else if l <= 3 && l > 0 && rd == 0 {
		// if the element is 0 < x < 3 and no dots are left
		return []string{s + "." + p}
	}

	var res []string
	var options []string

	// get the last biggest number (3-1 digits)
	max := int(math.Min(float64(l), 3))
	end := s[l-max:]

	// convert to number
	num, _ := strconv.Atoi(end)

	// check if number is valid and long enough
	if num <= 255 {
		for i := max; i > 0; i-- {
			if l >= i {
				options = append(options, RestoreIp(rd-1, s[l-i:], s[0:l-i])...)
			}
		}
	}

	fmt.Println(options)

	// append all possible numbers
	for i := 0; i < len(options); i++ {
		value := options[i]
		if p != "" {
			value += "." + p
		}
		res = append(res, value)
	}

	return res
}
