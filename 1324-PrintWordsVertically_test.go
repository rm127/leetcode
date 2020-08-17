package main

import (
	"reflect"
	"testing"
)

func Test_printVertically(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{"HOW ARE YOU"}, []string{"HAY", "ORO", "WEU"}},
		{"Example 2", args{"TO BE OR NOT TO BE"}, []string{"TBONTB", "OEROOE", "   T"}},
		{"Example 3", args{"CONTEST IS COMING"}, []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
		{"Trailing spaces", args{"NOT IS IS IS"}, []string{"NIII", "OSSS", "T"}},
		{"Trailing spaces rev", args{"IS IS IS NOT"}, []string{"IIIN", "SSSO", "   T"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printVertically(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
