package tree

import (
	"reflect"
	"testing"
)

func TestConvert2BiNode(t *testing.T) {
	type args struct {
		t *BSTree
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "10 nodes bst tree",
			args: args{t: NewBSTreeFromSortedArray([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})},
			want: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert2BiNode(tt.args.t.root)
			if !reflect.DeepEqual(serializeTreeNodeList(got), tt.want) {
				t.Errorf("Convert2BiNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
