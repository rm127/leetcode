package main

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"23", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", args{"2"}, []string{"a", "b", "c"}},
		{"[empty]", args{""}, []string{}},
		{"25", args{"25"}, []string{"aj", "ak", "al", "bj", "bk", "bl", "cj", "ck", "cl"}},
		{"247", args{"247"}, []string{"agp", "agq", "agr", "ags", "ahp", "ahq", "ahr", "ahs", "aip", "aiq", "air", "ais", "bgp", "bgq", "bgr", "bgs", "bhp", "bhq", "bhr", "bhs", "bip", "biq", "bir", "bis", "cgp", "cgq", "cgr", "cgs", "chp", "chq", "chr", "chs", "cip", "ciq", "cir", "cis"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LetterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LetterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
