package main

func GenerateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return []string{""}
	} else {
		for c := 0; c < n; c++ {
			for _, left := range GenerateParenthesis(c) {
				for _, right := range GenerateParenthesis(n - 1 - c) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
