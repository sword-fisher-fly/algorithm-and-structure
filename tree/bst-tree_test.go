package tree

import (
	"reflect"
	"testing"
)

func TestNewBSTreeFromSortedArray(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "3 items",
			args: args{arr: []any{1, 2, 3}},
			want: []any{1, 2, 3},
		},
		{
			name: "5 items",
			args: args{arr: []any{1, 2, 3, 4, 5}},
			want: []any{1, 2, 3, 4, 5},
		},
		{
			name: "10 items",
			args: args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBSTreeFromSortedArray(tt.args.arr)

			if got := tree.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBSTreeFromSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBSTreeFromSortedArrayByIteration(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "3 items",
			args: args{arr: []any{1, 2, 3}},
			want: []any{1, 2, 3},
		},
		{
			name: "5 items",
			args: args{arr: []any{1, 2, 3, 4, 5}},
			want: []any{1, 2, 3, 4, 5},
		},
		{
			name: "10 items",
			args: args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBSTreeFromSortedArrayByIteration(tt.args.arr)

			if got := tree.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBSTreeFromSortedArrayByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Insert(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "3 items",
			args: args{arr: []any{1, 2, 3}},
			want: []any{1, 2, 3},
		},
		{
			name: "5 items",
			args: args{arr: []any{1, 2, 3, 4, 5}},
			want: []any{1, 2, 3, 4, 5},
		},
		{
			name: "10 items",
			args: args{arr: []any{3, 6, 5, 7, 4, 1, 2, 9, 8, 10}},
			want: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &BSTree{}

			for _, val := range tt.args.arr {
				tr.Insert(val)
			}

			if got := tr.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_InsertByRecursive(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "3 items",
			args: args{arr: []any{1, 2, 3}},
			want: []any{1, 2, 3},
		},
		{
			name: "5 items",
			args: args{arr: []any{1, 2, 3, 4, 5}},
			want: []any{1, 2, 3, 4, 5},
		},
		{
			name: "10 items",
			args: args{arr: []any{3, 6, 5, 7, 4, 1, 2, 9, 8, 10}},
			want: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &BSTree{}

			for _, val := range tt.args.arr {
				tr.InsertByRecursive(val)
			}

			if got := tr.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTree_Search(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name   string
		args   args
		search any
		want   any
	}{
		{
			name:   "search item appears in 10 items",
			args:   args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			search: 7,
			want:   7,
		},
		{
			name:   "search item does not appear in 10 items",
			args:   args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			search: 12,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewBSTreeFromSortedArray(tt.args.arr)
			got := tr.Search(tt.search)
			// if got != nil {
			// 	t.Logf("got: %v, but want: %v", got.Val, tt.want)
			// }

			if got != nil && !reflect.DeepEqual(got.Val, tt.want) || got == nil && tt.want != nil {
				t.Errorf("failed to search %v from binary search tree", tt.search)
			}
		})

	}
}

func TestBSTree_SearchByRecursive(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name   string
		args   args
		search any
		want   any
	}{
		{
			name:   "search item appears in 10 items",
			args:   args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			search: 7,
			want:   7,
		},
		{
			name:   "search item does not appear in 10 items",
			args:   args{arr: []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			search: 12,
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewBSTreeFromSortedArray(tt.args.arr)

			got := tr.SearchByRecursive(tt.search)
			if got != nil && !reflect.DeepEqual(got.Val, tt.want) || got == nil && tt.want != nil {
				t.Errorf("BSTree.SearchByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
