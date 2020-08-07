package main

func ContainerWithMostWater(height []int) int {
	max := 0

	i, j := 0, len(height)-1

	for i < j {
		temp := (j - i) * Min(height[i], height[j])
		if temp > max {
			max = temp
		}

		if height[i] <= height[j] {
			i++
			continue
		}

		j--
	}
	return max
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
