package main

import "testing"

func TestRomanToInteger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[empty]", args{""}, 0},
		{"III", args{"III"}, 3},
		{"IV", args{"IV"}, 4},
		{"IX", args{"IX"}, 9},
		{"LVIII", args{"LVIII"}, 58},
		{"MCMXCIV", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInteger(tt.args.s); got != tt.want {
				t.Errorf("RomanToInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
