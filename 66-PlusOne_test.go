package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"123", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"4321", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"none", args{[]int{0}}, []int{1}},
		{"109", args{[]int{1, 0, 9}}, []int{1, 1, 0}},
		{"999", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
		{"9", args{[]int{9}}, []int{1, 0}},
		{"8999", args{[]int{8, 9, 9, 9}}, []int{9, 0, 0, 0}},
		{"989", args{[]int{9, 8, 9}}, []int{9, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
