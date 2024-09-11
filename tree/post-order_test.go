package tree

import (
	"reflect"
	"testing"
)

// postOrder: 4,5,2,6,7,3,1
//   1
//  / \
//  2  3
// /\  /\
// 4 5 6 7

func TestPostOrderStack(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want []any
	}{
		{
			name: "2-level",
			arr:  []any{1, 2, 3},
			want: []any{2, 3, 1},
		},
		{
			name: "3-level",
			arr:  []any{1, 2, 3, 4, 5, 6, 7},
			want: []any{4, 5, 2, 6, 7, 3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(tt.arr)

			if got := PostOrderStack(tree.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrderStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostOrderSingleStack(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want []any
	}{
		{
			name: "2-level",
			arr:  []any{1, 2, 3},
			want: []any{2, 3, 1},
		},
		{
			name: "3-level",
			arr:  []any{1, 2, 3, 4, 5, 6, 7},
			want: []any{4, 5, 2, 6, 7, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(tt.arr)

			if got := PostOrderSingleStack(tree.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrderSingleStack() = %v, want %v", got, tt.want)
			}

			treeII := NewTreeFromArray(tt.arr)
			if got := PostOrderStackII(treeII.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrderSingleStackII() = %v, want %v", got, tt.want)
			}
		})
	}
}
