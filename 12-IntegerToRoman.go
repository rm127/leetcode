package main

func IntegerToRoman(num int) string {
	letters := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}
	numbers := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	res := ""
	// keep last match, to avoid looking through numbers too big
	prevIndex := 0

	// while the input is not zero
	for num > 0 {
		for i := prevIndex; i < len(numbers); i++ {
			if num >= numbers[i] {
				prevIndex = i
				num = num - numbers[i]
				res = res + letters[i]
				break
			}
		}
	}

	return res
}
