package main

func ZigZagConversion(s string, numRows int) string {
	rows := make([]string, numRows)
	var res string

	i, idx := 0, 0
	growing := true
	for idx < len(s) {
		rows[i] += string(s[idx])
		if numRows > 1 {
			if i == numRows-1 {
				growing = false
			} else if i == 0 {
				growing = true
			}
			if growing {
				i++
			} else {
				i--
			}
		}
		idx++
	}
	for _, row := range rows {
		res += row
	}
	return res
}

/*

P   A   H   N		0, 4, 8, 12
A P L S I I G		1, 3, 5, 7, 9, 11
Y   I   R			2, 6, 10


P     I    N		0, 6, 12
A   L S  I G		1, 5, 7, 11
Y A   H R			2, 4, 8, 10
P     I				3, 9


P       I			0, 8
A     H R			1, 7, 9
Y   I   I			2, 6, 10
P L     N			3, 5, 11
A       G			4, 12

*/