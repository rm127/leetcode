package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "n2",
			args: args{2},
			want: []string{
				"()()",
				"(())",
			},
		},
		{
			name: "n3",
			args: args{3},
			want: []string{
				"()()()",
				"()(())",
				"(())()",
				"(()())",
				"((()))",
			},
		},
		{
			name: "n4",
			args: args{4},
			want: []string{
				"()()()()",
				"()()(())",
				"()(())()",
				"()(()())",
				"()((()))",
				"(())()()",
				"(())(())",
				"(()())()",
				"((()))()",
				"(()()())",
				"(()(()))",
				"((())())",
				"((()()))",
				"(((())))",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
