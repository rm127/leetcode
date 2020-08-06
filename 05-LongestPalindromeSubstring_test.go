package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"babad", args{"babad"}, "aba"},
		{"ssss", args{"ssss"}, "ssss"},
		{"asdfghjk", args{"asdfghjk"}, "k"},
		{"cbbd", args{"cbbd"}, "bb"},
		{"ccc", args{"ccc"}, "ccc"},
		{"abadd", args{"abadd"}, "aba"},
		{"aaabaaaa", args{"aaabaaaa"}, "aaabaaa"},
		{"aaaabaaa", args{"aaaabaaa"}, "aaabaaa"},
		{"xcyy", args{"xcyy"}, "yy"},
		{"tthj", args{"tthj"}, "tt"},
		{"rr", args{"rr"}, "rr"},
		{"wwwwww", args{"wwwwww"}, "wwwwww"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
