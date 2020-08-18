package main

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"digit", args{"0"}, true},
		{"decimal", args{" 0.1 "}, true},
		{"letters", args{"abc"}, false},
		{"space between", args{"1 a"}, false},
		{"exponent", args{"2e10"}, true},
		{"whitespace exp", args{" -90e3   "}, true},
		{"exp no exp", args{"1e"}, false},
		{"exp no lead", args{"e3"}, false},
		{"neg exp", args{"6e-1"}, true},
		{"incomplete exp", args{"6e-"}, false},
		{"decimal exp", args{"99e2.5"}, false},
		{"exp decimal lead", args{"53.5e93"}, true},
		{"double neg", args{" --6 "}, false},
		{"pos neg", args{"-+3"}, false},
		{"mixed letters", args{"95a54e53"}, false},
		{"just op", args{"-"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
