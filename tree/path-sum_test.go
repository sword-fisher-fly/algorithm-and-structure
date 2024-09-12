package tree

import (
	"reflect"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	type args struct {
		t   *Tree
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-level tree has path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 3,
			},
			want: true,
		},
		{
			name: "2-level tree has no path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 4,
			},
			want: false,
		},
		{
			name: "3-level tree has path sum",
			args: args{
				t:   level3FullTree,
				sum: 8,
			},
			want: true,
		},
		{
			name: "3-level tree has no path sum",
			args: args{
				t:   level3FullTree,
				sum: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSum(tt.args.t.root, tt.args.sum); got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPathSumSimple(t *testing.T) {
	type args struct {
		t   *Tree
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-level tree has path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 3,
			},
			want: true,
		},
		{
			name: "2-level tree has no path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 4,
			},
			want: false,
		},
		{
			name: "3-level tree has path sum",
			args: args{
				t:   level3FullTree,
				sum: 8,
			},
			want: true,
		},
		{
			name: "3-level tree has no path sum",
			args: args{
				t:   level3FullTree,
				sum: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSumSimple(tt.args.t.root, tt.args.sum); got != tt.want {
				t.Errorf("HasPathSumSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPathSumByIteration(t *testing.T) {
	type args struct {
		t   *Tree
		sum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-level tree has path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 3,
			},
			want: true,
		},
		{
			name: "2-level tree has no path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 2}),
				sum: 4,
			},
			want: false,
		},
		{
			name: "3-level tree has path sum",
			args: args{
				t:   level3FullTree, //{1, 2, 3, 4, 5, 6, 7}
				sum: 8,
			},
			want: true,
		},
		{
			name: "3-level tree has no path sum",
			args: args{
				t:   level3FullTree,
				sum: 20,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPathSumByIteration(tt.args.t.root, tt.args.sum); got != tt.want {
				t.Errorf("HasPathSumByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathSum(t *testing.T) {
	type args struct {
		t   *Tree
		sum int
	}
	tests := []struct {
		name string
		args args
		want [][]any
	}{
		{
			name: "3-level tree has path sum",
			args: args{
				t:   level3FullTree,
				sum: 8,
			},
			want: [][]any{{1, 2, 5}},
		},
		{
			name: "3-level tree has path sum",
			args: args{
				t:   level3FullTree,
				sum: 7,
			},
			want: [][]any{{1, 2, 4}},
		},
		{
			name: "random tree has two path sum",
			args: args{
				t:   NewTreeFromArray([]any{1, 4, 7, 5, 2, 2, -1}),
				sum: 10,
			},
			want: [][]any{{1, 4, 5}, {1, 7, 2}},
		},
		{
			name: "random tree has three path sum",
			args: args{
				t:   NewTreeFromArray([]any{5, 4, 8, 11, -1, 9, 4, 7, 2, -1, -1, -1, -1, 5, 1}),
				sum: 22,
			},
			want: [][]any{
				{5, 4, 11, 2},
				{5, 8, 9},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathSum(tt.args.t.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
