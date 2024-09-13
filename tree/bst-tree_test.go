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

func TestNewBSTreeFromLevelArray(t *testing.T) {
	type args struct {
		arr []any
	}
	tests := []struct {
		name        string
		args        args
		wantInOrder []any
	}{
		{
			name:        "2-level full binary search tree traversed in a level way",
			args:        args{arr: []any{2, 1, 3}},
			wantInOrder: []any{1, 2, 3},
		},
		//[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
		{
			name:        "any binary search tree traversed in a level way",
			args:        args{arr: []any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}},
			wantInOrder: []any{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBSTreeFromLevelArray(tt.args.arr)
			if !reflect.DeepEqual(got.InOrder(), tt.wantInOrder) {
				t.Errorf("NewBSTreeFromLevelArray() = %v, want %v", got, tt.wantInOrder)
			}
		})
	}
}
