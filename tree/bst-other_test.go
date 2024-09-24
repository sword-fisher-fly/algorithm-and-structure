package tree

import (
	"reflect"
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

func TestFindMode(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			//[1,2,3,4,5,5,6,7,8]
			name: "one mode in binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 5, 6, 7, 8}),
			},
			want: []any{5},
		},
		{
			name: "two modes in binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 4, 5, 5, 6, 7, 8}),
			},
			want: []any{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMode(tt.args.t.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindModeByIteration(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			//[1,2,3,4,5,5,6,7,8]
			name: "one mode in binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 5, 6, 7, 8}),
			},
			want: []any{5},
		},
		{
			name: "two modes in binary search tree",
			args: args{
				t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 4, 5, 5, 6, 7, 8}),
			},
			want: []any{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindModeByIteration(tt.args.t.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindModeByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimBSTreeByRecursive(t *testing.T) {
	type args struct {
		t    *BSTree
		low  int
		high int
	}
	tests := []struct {
		name        string
		args        args
		wantInOrder []any
	}{
		{
			name: "the value of root node is less than low",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  5,
				high: 7,
			},
			wantInOrder: []any{5, 6, 7},
		},
		{
			name: "the value of root node is between low and high",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  4,
				high: 7,
			},
			wantInOrder: []any{4, 5, 6, 7},
		},

		{
			name: "the value of root node is greater than high",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  7,
				high: 8,
			},
			wantInOrder: []any{7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimBSTreeByRecursive(tt.args.t.root, tt.args.low, tt.args.high); !reflect.DeepEqual(InOrderII(got), tt.wantInOrder) {
				t.Errorf("TrimBSTreeByRecursive() = %v, want %v", InOrderII(got), tt.wantInOrder)
			}
		})
	}
}

func TestTrimBSTreeByIteration(t *testing.T) {
	type args struct {
		t    *BSTree
		low  int
		high int
	}
	tests := []struct {
		name        string
		args        args
		wantInOrder []any
	}{
		{
			name: "the value of root node is less than low",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  5,
				high: 7,
			},
			wantInOrder: []any{5, 6, 7},
		},
		{
			name: "the value of root node is between low and high",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  4,
				high: 7,
			},
			wantInOrder: []any{4, 5, 6, 7},
		},

		{
			name: "the value of root node is greater than high",
			args: args{
				t:    NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				low:  7,
				high: 8,
			},
			wantInOrder: []any{7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimBSTreeByIteration(tt.args.t.root, tt.args.low, tt.args.high); !reflect.DeepEqual(InOrderII(got), tt.wantInOrder) {
				t.Errorf("TrimBSTreeByIteration() = %v, want %v", InOrderII(got), tt.wantInOrder)
			}
		})
	}
}

func TestGenerateBSTree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "only one node",
			args: args{n: 1},
			//
		},
		{
			name: "three node",
			args: args{n: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateBSTree(tt.args.n)

			for i := range got {
				t.Logf("Level traversal order: %v", LevelTraversal(got[i]))
			}
			//  !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GenerateBSTree() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestGenerateBSTreeCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3 nodes in a tree",
			args: args{n: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateBSTreeCount(tt.args.n); got != tt.want {
				t.Errorf("GenerateBSTreeCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
