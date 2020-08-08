package main

import "fmt"

func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trapped := 0

	for i := 0; i < len(height); i++ {
		fmt.Println(i, "trapped:", trapped)
		tempTrapped := 0
		for j := i + 1; j < len(height); j++ {
			fmt.Println("--> expanding: ", i, j, "temp: ", tempTrapped)
			if height[j] >= height[i] {
				fmt.Println("# reached end")
				trapped += tempTrapped
				i = j - 1
				break
			} else {
				tempTrapped += height[i] - height[j]
			}
		}
	}

	return trapped
}

/*

#
#       #
# # # # #

*/
