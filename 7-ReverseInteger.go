package main

import (
	"math"
)

func reverse(x int) int {
	rest := x
	size := int(math.Log10(math.Abs(float64(x))) + 1)
	res := 0
	for i := 1; i <= size; i++ {
		digit := rest % 10
		rest /= 10
		res = res * 10 + digit
		// limit for stupid 32-bit integer overflow
		if causesOverflow(res) {
			res = 0
			break
		}
	}
	return res
}

func causesOverflow(x int) bool {
	res := false
	if x > math.MaxInt32 || x < math.MinInt32 {
		res = true
	}
	return res
}