package main

import "testing"

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"Hello, my name is John"}, 5},
		{"Only one", args{"Nothing"}, 1},
		{"Many", args{"a a a a a a a a"}, 8},
		{"Empty", args{""}, 0},
		{"Trickster", args{"                "}, 0},
		{"Strip", args{" asjdnasd "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
