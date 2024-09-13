package tree

import (
	"reflect"
	"testing"
)

func TestAccumulateTreeNodeByRecursive(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name        string
		args        args
		wantInorder []any
	}{
		{
			name:        "2-level full tree",
			args:        args{t: NewBSTreeFromSortedArray([]any{1, 2, 3})},
			wantInorder: []any{6, 5, 3},
		},
		//输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
		//输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
		{
			name:        "random binary search tree",
			args:        args{t: NewBSTreeFromSortedArray([]any{0, 1, 2, 3, 4, 5, 6, 7, 8})},
			wantInorder: []any{36, 36, 35, 33, 30, 26, 21, 15, 8},
		},
		{
			name:        "construct binary search tree from level array",
			args:        args{t: NewBSTreeFromLevelArray([]any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8})},
			wantInorder: []any{36, 36, 35, 33, 30, 26, 21, 15, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccumulateTreeNodeByRecursive(tt.args.t.root); !reflect.DeepEqual(tt.args.t.InOrder(), tt.wantInorder) {
				t.Errorf("AccumulateTreeNodeByRecursive() = %v, want %v", got, tt.wantInorder)
			}
		})
	}
}
