package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"basic", args{"abcabcbb"}, 3},
		{"one", args{"e"}, 1},
		{"same", args{"yyyyyy"}, 1},
		{"reverse", args{"cba"}, 3},
		{"pwwkew", args{"pwwkew"}, 3},
		{"ridiculous", args{"qwqwqwqwqwqwqwqwqwqabcde"}, 7},
		{"asdfgghjkl", args{"asdfgghjkl"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
