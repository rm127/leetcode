package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fl", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"[none]", args{[]string{"dog", "racecar", "car"}}, ""},
		{"[empty]", args{[]string{""}}, ""},
		{"a", args{[]string{"a"}}, "a"},
		{"aaa", args{[]string{"aaa", "aa", "aaa"}}, "aa"},
		{"s", args{[]string{"something", "sand", "sahara"}}, "s"},
		{"all", args{[]string{"all", "all", "all"}}, "all"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
