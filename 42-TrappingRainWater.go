package main

import "fmt"

func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trapped := 0
	var currentlvl int
	tempTrapped := 0

	for _, lvl := range height {
		fmt.Println("lvl: ", lvl, "curlvl: ", currentlvl, "temp: ", tempTrapped, "trapped:", trapped)
		if lvl < currentlvl {
			//fmt.Println("adding to temp: ", currentlvl-lvl)
			tempTrapped += currentlvl - lvl
		}
		if lvl >= currentlvl {
			currentlvl = lvl
			//fmt.Println("moving to final: ", tempTrapped)
			trapped += tempTrapped
			tempTrapped = 0
		}
	}

	return trapped
}
