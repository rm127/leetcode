package main

func TrappingRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	trapped := 0

	for i := 0; i < len(height); i++ {
		tempTrapped := 0
		origHeight := height[i]
		done := false
		// loop for column height
		for !done {
			tempTrapped = 0
			// increase in size to the right until column of
			// same height or higher is reached
			for j := i + 1; j < len(height); j++ {
				// if same height or heigher is reached = done
				if height[j] >= origHeight {
					trapped += tempTrapped
					// set index to end of dip
					i = j - 1
					done = true
					break
				} else {
					tempTrapped += origHeight - height[j]
				}
			}
			// as long as the columns temp height isn't at the
			//// bottom => decrease in height and run again
			if origHeight > 0 {
				origHeight--
			} else {
				done = true
			}
		}
	}
	return trapped
}
