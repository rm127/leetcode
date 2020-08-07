package main

func ContainerWithMostWater(height []int) int {
	max := 0

	for i := 0; i < len(height)-1; i++ {
		if (len(height)-i-1)*height[i] <= max {
			continue
		}
		for j := i + 1; j < len(height); j++ {
			temp := (j - i) * Min(height[i], height[j])
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
