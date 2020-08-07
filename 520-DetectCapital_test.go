package main

import "testing"

func TestDetectCapital(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"all uppercase", args{"USA"}, true},
		{"FlaG", args{"FlaG"}, false},
		{"all downcase", args{"something"}, true},
		{"G", args{"G"}, true},
		{"s", args{"s"}, true},
		{"mL", args{"mL"}, false},
		{"Ml", args{"Ml"}, true},
		{"Leetcode", args{"Leetcode"}, true},
		{"ffffffffffffffffffffF", args{"ffffffffffffffffffffF"}, false},
		{"FFFFFFFFFFFFFFFFFFFFf", args{"FFFFFFFFFFFFFFFFFFFFf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectCapital(tt.args.word); got != tt.want {
				t.Errorf("DetectCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
