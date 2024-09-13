package tree

import (
	"reflect"
	"testing"
)

// preOrder: 1,2,4,5,3
//
//	 1
//	/ \
//	2  3
//
// /\
// 4 5
func TestPreOrderStack(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want []any
	}{
		{
			name: "2-level tree",
			arr:  []any{1, 2, 3},
			want: []any{1, 2, 3},
		},
		{
			name: "3-level tree",
			arr:  []any{1, 2, 3, 4, 5},
			want: []any{1, 2, 4, 5, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(tt.arr)

			if got := PreOrderStack(tree.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrderStack() = %v, want %v", got, tt.want)
			}

			treeII := NewTreeFromArray(tt.arr)
			if got := PreOrderStackII(treeII.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrderStackII() = %v, want %v", got, tt.want)
			}

			// treeIII := NewTreeFromArray(tt.arr)
			// if got := PreOrderStackIII(treeIII.root); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("PreOrderStackIII() = %v, want %v", got, tt.want)
			// }
		})
	}
}
