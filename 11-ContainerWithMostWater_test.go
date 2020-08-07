package main

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"Custom", args{[]int{1, 2, 3, 4, 5}}, 6},
		{"Equal", args{[]int{3, 3, 3, 3}}, 9},
		{"New", args{[]int{1, 3, 2, 5, 25, 24, 5}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainerWithMostWater(tt.args.height); got != tt.want {
				t.Errorf("ContainerWithMostWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
