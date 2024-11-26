package tree

import (
	"reflect"
	"testing"
)

func TestConstructTreeFromPreOrderAndInOrder(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name          string
		args          args
		wantPostOrder []any
	}{
		{
			name: "2-level full tree",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{2, 1, 3},
			},
			wantPostOrder: []any{2, 3, 1},
		},
		{
			name: "3-level full tree",
			args: args{
				preorder: []int{1, 2, 4, 5, 3, 6, 7},
				inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			},
			wantPostOrder: []any{4, 5, 2, 6, 7, 3, 1},
		},
		//[1,2,4,5,6,3],[5,4,6,2,1,3]
		{
			name: "case 3",
			args: args{
				preorder: []int{1, 2, 4, 5, 6, 3},
				inorder:  []int{5, 4, 6, 2, 1, 3},
			},
			wantPostOrder: []any{5, 6, 4, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructTreeFromPreOrderAndInOrder(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(PostOrderII(got), tt.wantPostOrder) {
				t.Errorf("ConstructTreeFromPreOrderAndInOrder() = %v, want %v", got, tt.wantPostOrder)
			}
		})
	}
}

func TestConstructTreeFromInOrderAndPostOrder(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name         string
		args         args
		wantPreOrder []any
	}{
		{
			name: "2-level full tree",
			args: args{
				inorder:   []int{2, 1, 3},
				postorder: []int{2, 3, 1},
			},
			wantPreOrder: []any{1, 2, 3},
		},
		{
			name: "3-level full tree",
			args: args{
				inorder:   []int{4, 2, 5, 1, 6, 3, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			wantPreOrder: []any{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructTreeFromInOrderAndPostOrder(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(PreOrderStackII(got), tt.wantPreOrder) {
				t.Errorf("ConstructTreeFromInOrderAndPostOrder() = %v, want preorder %v", got, tt.wantPreOrder)
			}
		})
	}
}
