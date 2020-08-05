package main

import "testing"

func TestZigZagConversion(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"PAYPALISHIRING-3", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING-4", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"AB-4", args{"AB", 1}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZigZagConversion(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("ZigZagConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}