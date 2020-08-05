package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"same value", args{[]int{3,3}, 6},[]int{0,1}},
		{"1 to 6", args{[]int{1,2,3,4,5,6}, 7},[]int{2,3}},
		{"same zero", args{[]int{0,0}, 0},[]int{0,1}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}