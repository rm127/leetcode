package main

import "testing"

func TestTrappingRainWater(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Cradle", args{[]int{4, 1, 4}}, 3},
		{"Simple", args{[]int{1, 0, 2}}, 1},
		{"Example", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"Big cradle", args{[]int{3, 0, 0, 0, 3}}, 9},
		{"Hill", args{[]int{1, 2, 3, 2, 1}}, 0},
		{"Annoying", args{[]int{1, 0, 3, 1, 2, 1, 3, 0, 1}}, 7},
		{"Tower", args{[]int{0, 0, 0, 4, 0, 0}}, 0},
		{"Slope", args{[]int{9, 7, 8, 6, 7, 5, 6, 4, 5, 3, 4, 2, 3, 1, 2}}, 7},
		{"Reservoir", args{[]int{3, 1, 1, 1, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrappingRainWater(tt.args.height); got != tt.want {
				t.Errorf("TrappingRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
