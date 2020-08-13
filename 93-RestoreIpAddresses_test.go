package main

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example", args{"25525511135"}, []string{"255.255.11.135", "255.255.111.35"}},
		{"Same", args{"11111111111"}, []string{"111.111.111.11", "111.111.11.111", "111.11.111.111", "11.111.111.111"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
