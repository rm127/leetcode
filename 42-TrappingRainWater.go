package main

import "fmt"

func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trapped := 0
	currentlvl := height[0]

	for i, lvl := range height {
		fmt.Println("lvl: ", lvl, "curlvl: ", currentlvl, "trapped:", trapped)
		if lvl < currentlvl {
			currentlvl = lvl
		} else if lvl > currentlvl {
			fmt.Println("--> LOOKBACK")
			tempTrapped := 0
			for j := i - 1; j >= 0; j-- {
				fmt.Println("|--> lvl: ", lvl, "elem: ", height[j], "temp: ", tempTrapped)
				if height[j] >= lvl {
					trapped += tempTrapped
					break
				} else {
					tempTrapped += lvl - height[j]
				}
			}
		}
	}

	return trapped
}
