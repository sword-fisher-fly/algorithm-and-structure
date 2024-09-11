package tree

import (
	"testing"
)

func convertByteSliceToString(arr []any) string {
	bs := make([]byte, len(arr))
	for i := range arr {
		bs[i] = byte(arr[i].(int32))
	}

	return string(bs)
}

func TestAlphabetTree(t *testing.T) {

	tests := []struct {
		name      string
		arr       []any
		preOrder  string
		inOrder   string
		postOrder string
	}{
		{
			name:      "alphabet tree",
			arr:       alphabetArr,
			preOrder:  "ABDHIEJCFKG",
			inOrder:   "HDIBEJAFKCG",
			postOrder: "HIDJEBKFGCA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(alphabetArr)

			if preOrder := tree.PreOrder(); convertByteSliceToString(preOrder) != tt.preOrder {
				t.Logf("preOrder(alphabet): got %v, want %v", convertByteSliceToString(preOrder), tt.preOrder)

			}

			if inOrder := tree.InOrder(); convertByteSliceToString(inOrder) != tt.inOrder {
				t.Logf("inOrder(alphabet): got %v, want %v", convertByteSliceToString(inOrder), tt.inOrder)
			}

			if postOrder := tree.PostOrder(); convertByteSliceToString(postOrder) != tt.postOrder {
				t.Logf("postOrder(alphabet): got %v, want %v", convertByteSliceToString(postOrder), tt.postOrder)
			}
		})
	}
}
