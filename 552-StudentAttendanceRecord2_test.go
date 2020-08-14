package main

import "testing"

func Test_checkRecord2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{2}, 8},
		{"1", args{1}, 3},
		{"3", args{3}, 19},
		{"4", args{4}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord2(tt.args.n); got != tt.want {
				t.Errorf("checkRecord2() = %v, want %v", got, tt.want)
			}
		})
	}
}
