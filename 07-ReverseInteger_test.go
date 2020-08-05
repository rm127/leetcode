package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"123", args{123}, 321},
		{"584", args{584}, 485},
		{"1234", args{1234}, 4321},
		{"10", args{10}, 1},
		{"87654", args{87654}, 45678},
		{"-123", args{-123}, -321},
		{"1534236469", args{1534236469}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
