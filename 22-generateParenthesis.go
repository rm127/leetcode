package main

func generateParenthesis(n int) []string {
	var res []string

	if n == 0 {
		return []string{""}
	} else {
		for c := 0; c < n; c++ {
			for _, left := range generateParenthesis(c) {
				for _, right := range generateParenthesis(n - 1 - c) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
