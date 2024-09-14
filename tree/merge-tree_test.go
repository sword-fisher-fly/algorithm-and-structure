package tree

import (
	"reflect"
	"testing"
)

// Tree1:
//
//	   4
//	   /\
//	  3  2
//	 /    \
//	5      7
//	/\
//
// 8  9
// [4,3,2,5,-1,-1,7,8,9]
// Tree2:
//
//	    1
//	  /  \
//	  5   6
//	 /\   /
//	10 11 8
//	      /
//	     12
//
// [1,5,6,10,11,8,-1,-1,-1,-1,-1,12]
// merge Tree1:
//
//	    5
//	  /   \
//	  8   8
//	 /\   /\
//	15 11 8 7
//	/\    /
//
// 8  9   12
//
//	[5,8,8,15,11,8,7,8,9,-1,-1,12]
//
// preOrder: [5,8,15,8,9,11,8,8,12,7]
func TestMergeTrees(t *testing.T) {
	type args struct {
		t1 *Tree
		t2 *Tree
	}
	tests := []struct {
		name         string
		args         args
		wantPreOrder []any
	}{
		{
			name: "the same tree",
			args: args{
				t1: level2FullTree, // 1,2,3
				t2: level2FullTree,
			},
			wantPreOrder: []any{2, 4, 6},
		},
		{
			name: "different tree",
			args: args{
				t1: NewTreeFromArray([]any{4, 3, 2, 5, -1, -1, 7, 8, 9}),
				t2: NewTreeFromArray([]any{1, 5, 6, 10, 11, 8, -1, -1, -1, -1, -1, 12}),
			},
			wantPreOrder: []any{5, 8, 15, 8, 9, 11, 8, 8, 12, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeTrees(tt.args.t1.root, tt.args.t2.root)
			if !reflect.DeepEqual(PreOrderStack(got), tt.wantPreOrder) {
				t.Errorf("MergeTrees() = %v, want preorder: %v", PreOrderStack(got), tt.wantPreOrder)
			}
		})
	}
}

func TestMergeTreesByIteration(t *testing.T) {
	type args struct {
		t1 *Tree
		t2 *Tree
	}
	tests := []struct {
		name         string
		args         args
		wantPreOrder []any
	}{
		{
			name: "the same tree",
			args: args{
				t1: level2FullTree, // 1,2,3
				t2: level2FullTree,
			},
			wantPreOrder: []any{2, 4, 6},
		},
		{
			name: "different tree",
			args: args{
				t1: NewTreeFromArray([]any{4, 3, 2, 5, -1, -1, 7, 8, 9}),
				t2: NewTreeFromArray([]any{1, 5, 6, 10, 11, 8, -1, -1, -1, -1, -1, 12}),
			},
			wantPreOrder: []any{5, 8, 15, 8, 9, 11, 8, 8, 12, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTreesByIteration(tt.args.t1.root, tt.args.t2.root); !reflect.DeepEqual(PreOrderStackII(got), tt.wantPreOrder) {
				t.Errorf("MergeTreesByIteration() = %v, want preorder: %v", PreOrderStackII(got), tt.wantPreOrder)
			}
		})
	}
}
