package tree

import (
	"reflect"
	"testing"
)

func TestTreePathFromRootToLeaf(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2-level tree",
			args: args{
				t: level2FullTree,
			},
			want: []string{"1->2", "1->3"},
		},
		{
			name: "3-level tree",
			args: args{
				t: level3FullTree,
			},
			want: []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"},
		},
		// 1, 2, 3, 4, 5, 6, 7}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreePathFromRootToLeaf(tt.args.t.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreePathFromRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreePathFromRootToLeafII(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2-level tree",
			args: args{
				t: NewTreeFromArray([]any{1, 2, 3}),
			},
			want: []string{"1->2", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreePathFromRootToLeafII(tt.args.t.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreePathFromRootToLeafII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreePathFromRootToLeafByIteration(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2-level tree",
			args: args{
				t: level2FullTree,
			},
			want: []string{"1->2", "1->3"},
		},
		{
			name: "3-level tree",
			args: args{
				t: level3FullTree,
			},
			want: []string{"1->2->4", "1->2->5", "1->3->6", "1->3->7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreePathFromRootToLeafByIteration(tt.args.t.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreePathFromRootToLeafByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
