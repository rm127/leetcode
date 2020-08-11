package main

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{
			&TreeNode{1,
				&TreeNode{2,
					&TreeNode{4,
						nil, nil},
					&TreeNode{5,
						nil, nil},
				}, &TreeNode{3, &TreeNode{6, nil, nil}, nil},
			},
		}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
