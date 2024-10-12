package tree

import (
	"testing"
)

func TestFindTwoSum(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 11,
			},
			want: true,
		},
		{
			name: "not found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTwoSum(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("FindTwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTwoSumII(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 11,
			},
			want: true,
		},
		{
			name: "not found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTwoSumII(tt.args.root, tt.args.target); got != tt.want {
				t.Logf("inorder bst: %v", InOrderStack(tt.args.root))
				t.Errorf("FindTwoSumII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTwoSumAdjacent(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 11,
			},
			want: true,
		},
		{
			name: "not found the sum of two nodes in the tree",
			args: args{
				root:   NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Root(),
				target: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTwoSumAdjacent(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("FindTwoSumAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
