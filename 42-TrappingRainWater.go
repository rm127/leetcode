package main

import "fmt"

func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trapped := 0
	var currentlvl int

	for i := 0; i < len(height); i++ {
		lvl := height[i]
		fmt.Println("lvl: ", lvl, "curlvl: ", currentlvl, "trapped:", trapped)
		overflow := false
		//tempTrapped := 0
		l, r := i, i
		max := lvl
		tempTrapped := 0
		for !overflow {
			fmt.Println("l: ", l, ", r: ", r, height[l:r+1], ", max: ", max, ", temp: ", tempTrapped)

			if l < i && r > i {
				tempTrapped = (r - l - 1) * max
			}

			if lvl > height[l] || lvl > height[r] {
				overflow = true
			}

			if l > 0 {
				l--
			} else if r < len(height)-1 {
				r++
			} else {
				// end reached = overflow
				fmt.Println("## end overflow")
				overflow = true
			}
		}
		trapped += tempTrapped
	}
	return trapped
}

/*

For each column
-> check all adjacent columns
-> IF a column is lower than the previous, then overflow
-> ELSE move up a row
*/
