package tree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only one root node",
			args: args{t: level1FullTree},
			want: 1,
		},
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: 2,
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepth(tt.args.t.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepthByBackTracking(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only one root node",
			args: args{t: level1FullTree},
			want: 1,
		},
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: 2,
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: 3,
		},
		{
			name: "4-level symmetric tree",
			args: args{t: level4SymmetricTree},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepthByBackTracing(tt.args.t.root); got != tt.want {
				t.Errorf("MaxDepthByBackTracking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxDepthByBFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only one root node",
			args: args{t: level1FullTree},
			want: 1,
		},
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: 2,
		},
		{
			name: "3-level tree",
			args: args{t: level3FullTree},
			want: 3,
		},
		{
			name: "4-level symmetric tree",
			args: args{t: level4SymmetricTree},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepthByBFS(tt.args.t.root); got != tt.want {
				t.Errorf("MaxDepthByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
