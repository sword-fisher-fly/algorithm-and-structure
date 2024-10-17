package recursive

import (
	"reflect"
	"testing"
)

func TestConstructBSTreeFromSortedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name        string
		args        args
		wantInOrder []interface{}
	}{
		{
			name:        "single number",
			args:        args{nums: []int{1}},
			wantInOrder: []interface{}{1},
		},
		{
			name:        "10 items with 5 consecutive nums",
			args:        args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			wantInOrder: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructBSTreeFromSortedArray(tt.args.nums); !reflect.DeepEqual(InOrderTraverse(got.root), tt.wantInOrder) {
				t.Errorf("ConstructBSTreeFromSortedArray() = %v, want %v", InOrderTraverse(got.root), tt.wantInOrder)
			}
		})
	}
}

func TestLevelTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name      string
		args      args
		wantLevel [][]interface{}
	}{
		{
			name:      "10 items with 5 consecutive nums",
			args:      args{root: ConstructBSTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).root},
			wantLevel: [][]interface{}{{5}, {2, 8}, {1, 3, 6, 9}, {4, 7, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelTraversal(tt.args.root); !reflect.DeepEqual(got, tt.wantLevel) {
				t.Errorf("LevelTraversal() = %v, want %v", got, tt.wantLevel)
			}
		})
	}
}
