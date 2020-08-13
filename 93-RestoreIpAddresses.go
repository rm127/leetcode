package main

import (
	"fmt"
	"math"
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	// an Ip address requires at least 4 digits
	if len(s) < 4 {
		return []string{}
	}
	return RestoreIp(3, "", s)
}

func RestoreIp(rd int, p string, s string) []string {
	var res []string
	var options []string

	l := len(s)

	if l > 3 && rd == 0 {
		// if there are no dots left and string is too long
		return []string{}
	} else if l <= 3 && l > 0 && rd == 0 {
		// if the element is 0 < x < 3 and no dots are left
		return []string{s + "." + p}
	}

	// get the last biggest number (3-1 digits)

	end := s[l-int(math.Min(float64(l), 3)):]

	num, _ := strconv.Atoi(end)

	fmt.Println(int(math.Min(float64(l), 3)), num)

	// check if number is valid and long enough
	if num <= 255 && l > 2 {
		options = append(options, RestoreIp(rd-1, s[l-3:], s[0:l-3])...)
	}
	if l > 1 {
		options = append(options, RestoreIp(rd-1, s[l-2:], s[0:l-2])...)
	}
	if l > 0 {
		options = append(options, RestoreIp(rd-1, s[l-1:], s[0:l-1])...)
	}

	for i := 0; i < len(options); i++ {
		value := options[i]
		if p != "" {
			value += "." + p
		}
		res = append(res, value)
	}

	return res
}
