package tree

import (
	"reflect"
	"testing"
)

// inOrder: 4,2,5,1,6,3,7
// inOrder: 8,4,9,2,10,5,11,1,12,6,13,3,14,7,15
//
//	  1
//	 / \
//	2  3
//	/\ /\
//
// 4    5     6      7
// /\   /\    /\    /\
// 8 9 10 11 12 13 14 15
func TestInOrderStack(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want []any
	}{
		{
			name: "2-level tree",
			arr:  []any{1, 2, 3},
			want: []any{2, 1, 3},
		},
		{
			name: "3-level tree",
			arr:  []any{1, 2, 3, 4, 5, 6, 7},
			want: []any{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "4-level tree",
			arr:  []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: []any{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(tt.arr)
			if got := InOrderStack(tree.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderStack() = %v, want %v", got, tt.want)
			}

			treeII := NewTreeFromArray(tt.arr)
			if got := InOrderStackII(treeII.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrderStackII() = %v, want %v", got, tt.want)
			}
		})
	}
}
