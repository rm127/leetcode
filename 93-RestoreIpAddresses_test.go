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
		{"Same", args{"11111111111"}, []string{"11.111.111.111", "111.11.111.111", "111.111.11.111", "111.111.111.11"}},
		{"Empty", args{""}, []string{}},
		{"Four", args{"1111"}, []string{"1.1.1.1"}},
		{"Four zero", args{"0000"}, []string{"0.0.0.0"}},
		{"Strip zeros", args{"010010"}, []string{"0.10.0.10", "0.100.1.0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
