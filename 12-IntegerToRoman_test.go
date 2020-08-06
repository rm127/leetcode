package main

import "testing"

func TestIntegerToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{0}, ""},
		{"3", args{3}, "III"},
		{"4", args{4}, "IV"},
		{"9", args{9}, "IX"},
		{"58", args{58}, "LVIII"},
		{"1994", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntegerToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntegerToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
