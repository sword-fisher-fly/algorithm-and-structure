package tree

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			},
			want: true,
		},
		{
			name: "invalid binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 6, 5, 7, 8, 9, 10, 11}),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBST(tt.args.t.root); got != tt.want {
				t.Log(tt.args.t.InOrder())

				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidBSTByIteration(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			},
			want: true,
		},
		{
			name: "invalid binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 6, 5, 7, 8, 9, 10, 11}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBSTByIteration(tt.args.t.root); got != tt.want {
				t.Errorf("IsValidBSTByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
