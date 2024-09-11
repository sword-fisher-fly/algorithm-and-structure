package tree

import (
	"reflect"
	"testing"
)

func TestLevelTraverse(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want [][]any
	}{
		{
			name: "1-level tree",
			arr:  level1Arr,
			want: [][]any{{1}},
		},
		{
			name: "2-level tree",
			arr:  level2Arr,
			want: [][]any{
				{1},
				{2, 3},
			},
		},
		{
			name: "3-level tree",
			arr:  level3Arr,
			want: [][]any{
				{1},
				{2, 3},
				{4, 5, 6, 7},
			},
		},
		{
			name: "4-level tree",
			arr:  level4Arr,
			want: [][]any{
				{1},
				{2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11, 12, 13, 14, 15},
			},
		},
		{
			name: "4-level random tree",
			arr:  level4RandomArr,
			want: [][]any{
				{1},
				{2, 6},
				{3, 4},
				{5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTreeFromArray(tt.arr)

			if got := LevelTraversal(tree.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
