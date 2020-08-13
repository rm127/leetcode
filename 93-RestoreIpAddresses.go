package main

import (
	"fmt"
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	return IpRestore("", s)
}

func IpRestore(prefix string, s string) []string {
	var res []string
	if len(s) < 3 {
		return []string{s}
	}

	var options []string

	current := s[0:3]

	convInt, _ := strconv.Atoi(current)

	fmt.Println("curr: ", convInt, s)

	if convInt > 255 {
		current = s[0:2]
		options = append(options, IpRestore(s[0:2], s[2:])...)
	} else {
		options = append(options, IpRestore(s[0:3], s[3:])...)
	}

	fmt.Println("opts: ", options)
	for i := 0; i < len(options); i++ {
		res = append(res, prefix+"."+options[i])
	}
	return res
}
