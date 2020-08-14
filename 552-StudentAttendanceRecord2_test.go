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
		{"1", args{1}, 3},
		{"2 (Example 1)", args{2}, 8},
		{"3", args{3}, 19},
		{"4", args{4}, 43},
		{"5", args{5}, 94},
		{"6", args{6}, 200},
		{"7", args{7}, 418},
		{"8", args{8}, 861},
		{"9", args{9}, 1753},
		{"10", args{10}, 3536},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord2(tt.args.n); got != tt.want {
				t.Errorf("checkRecord2() = %v, want %v", got, tt.want)
			}
		})
	}
}
