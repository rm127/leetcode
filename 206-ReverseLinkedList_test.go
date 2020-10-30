package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Example 1", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		{"Example 2", args{nil}, nil},
		{"Example 3", args{&ListNode{0, nil}}, &ListNode{0, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
