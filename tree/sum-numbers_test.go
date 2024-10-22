package tree

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: NewTreeFromArray([]any{1, 2, 3}).Root(),
			},
			want: 25,
		},
		{
			name: "case 2",
			args: args{
				root: NewTreeFromArray([]any{4, 9, 0, 5, 1}).Root(),
			},
			want: 1026,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumbers(tt.args.root); got != tt.want {
				t.Errorf("SumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumNumbersByLevelTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				root: NewTreeFromArray([]any{1, 2, 3}).Root(),
			},
			want: 25,
		},
		{
			name: "case 2",
			args: args{
				root: NewTreeFromArray([]any{4, 9, 0, 5, 1}).Root(),
			},
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumbersByLevelTraversal(tt.args.root); got != tt.want {
				t.Errorf("SumNumbersByLevelTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
