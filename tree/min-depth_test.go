package tree

import (
	"testing"
)

func TestMinDepth(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2-level full tree",
			args: args{
				t: level2FullTree,
			},
			want: 2,
		},
		{
			name: "4-level symmetric tree",
			args: args{
				t: level4SymmetricTree,
			},
			want: 3,
		},
		{
			name: "4-level non-symmetric tree",
			args: args{
				t: level4RandomTree,
			},
			want: 2,
		},
		// {1, 2, 6, 3, 4, -1 - 1, 5}
		//   1
		//   /\
		//  2  6
		// /\  /\
		// 3 4 -1 -1
		// /
		// 5
		{
			name: "4-level non-symmetric tree",
			args: args{
				t: NewTreeFromArray([]any{1, 2, -1, 3, 4, -1, -1, 5}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDepth(tt.args.t.root); got != tt.want {
				t.Errorf("MinDepth() = %v, want %v", got, tt.want)
				t.Errorf("level traversal tree: %v", tt.args.t.LevelTraversal())
			}

			if got := MinDepthII(tt.args.t.root); got != tt.want {
				t.Errorf("MinDepthII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinDepthByBFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2-level full tree",
			args: args{
				t: level2FullTree,
			},
			want: 2,
		},
		{
			name: "4-level symmetric tree",
			args: args{
				t: level4SymmetricTree,
			},
			want: 3,
		},
		{
			name: "4-level non-symmetric tree",
			args: args{
				t: level4RandomTree,
			},
			want: 2,
		},
		// {1, 2, 6, 3, 4, -1 - 1, 5}
		//   1
		//   /\
		//  2  6
		// /\  /\
		// 3 4 -1 -1
		// /
		// 5
		{
			name: "4-level non-symmetric tree",
			args: args{
				t: NewTreeFromArray([]any{1, 2, -1, 3, 4, -1, -1, 5}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinDepthByBFS(tt.args.t.root); got != tt.want {
				t.Errorf("MinDepthByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
