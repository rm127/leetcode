package main

import "testing"

func TestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{"()"}, want: true},
		{name: "2", args: args{"()[]{}"}, want: true},
		{name: "3", args: args{"(]"}, want: false},
		{name: "4", args: args{"([)]"}, want: false},
		{name: "5", args: args{"{[]}"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
