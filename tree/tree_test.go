package tree

import (
	"reflect"
	"testing"
)

// preOrder:  1 2 4 5 3 6
// inOrder:   4 2 5 1 6 3
// postOrder: 4 5 2 6 3 1
func TestTreeTraversal(t *testing.T) {
	testcases := []struct {
		name      string
		input     []any
		preOrder  []any
		inOrder   []any
		postOrder []any
		empty     bool
	}{
		{
			name:      "no root",
			input:     []any{},
			preOrder:  []any{},
			inOrder:   []any{},
			postOrder: []any{},
			empty:     true,
		},
		{
			name:      "only one root",
			input:     []any{1},
			preOrder:  []any{1},
			inOrder:   []any{1},
			postOrder: []any{1},
		},
		{
			name:      "3-level root",
			input:     []any{1, 2, 3, 4, 5, 6, -1},
			preOrder:  []any{1, 2, 4, 5, 3, 6},
			inOrder:   []any{4, 2, 5, 1, 6, 3},
			postOrder: []any{4, 5, 2, 6, 3, 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewTreeFromArray(tc.input)

			if !tree.Empty() {
				t.Logf("root node: %v", tree.Root().Val)
			}

			preOrderArr := tree.PreOrder()
			if !tree.Empty() && !reflect.DeepEqual(preOrderArr, tc.preOrder) {
				t.Errorf("preOrder got: %v, want: %v", preOrderArr, tc.preOrder)
			}

			inOrderArr := tree.InOrder()
			if !tree.Empty() && !reflect.DeepEqual(inOrderArr, tc.inOrder) {
				t.Errorf("inOrder got: %v, want: %v", inOrderArr, tc.inOrder)
			}

			postOrderArr := tree.PostOrder()
			if !tree.Empty() && !reflect.DeepEqual(postOrderArr, tc.postOrder) {
				t.Errorf("postOrder got: %v, want: %v", postOrderArr, tc.postOrder)
			}
		})
	}
}
