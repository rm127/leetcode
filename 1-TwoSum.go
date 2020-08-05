package main

func main() {
}

func TwoSum(nums []int, target int) []int {
	var res []int
	lookup := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if val, ok := lookup[complement]; ok && val != i {
			res = []int{val, i}
			break
		}
		lookup[num] = i
	}
	return res
}