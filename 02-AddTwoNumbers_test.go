package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"single", args{
			&ListNode{1, nil},
			&ListNode{2, nil},
		}, &ListNode{3, nil}},
		{"simple", args{
			&ListNode{1, &ListNode{2, nil}},
			&ListNode{3, &ListNode{4, nil}},
		}, &ListNode{4, &ListNode{6, nil}}},
		{"zero", args{
			&ListNode{0, nil},
			&ListNode{0, nil},
		}, &ListNode{0, nil}},
		{"overlap", args{
			&ListNode{9, nil},
			&ListNode{9, nil},
		}, &ListNode{8, &ListNode{1, nil}}},
		{"large", args{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
