package tree

import (
	"reflect"
	"testing"
)

// 1, 2, 3, 4, 5, 6, 7
//
//	  1             1
//	 /\             /\
//	2  3     ->    3  2
//
// /\ / \         /\   /\
// 4 5 6 7       7  6  5 4
//
//	1,3,5,4,2,7,6
func TestInvertTree(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: []any{1, 3, 2}, //preOrder
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: []any{1, 3, 7, 6, 2, 5, 4}, //preOrder
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InvertTree(tt.args.t.root)

			if got := tt.args.t.PreOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvertTreeByDFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: []any{1, 3, 2}, //preOrder
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: []any{1, 3, 7, 6, 2, 5, 4}, //preOrder
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InvertTreeByDFS(tt.args.t.root)

			if got := tt.args.t.PreOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTreeByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvertTreeByBFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: []any{1, 3, 2}, //preOrder
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: []any{1, 3, 7, 6, 2, 5, 4}, //preOrder
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InvertTreeByBFS(tt.args.t.root)

			if got := tt.args.t.PreOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTreeByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
