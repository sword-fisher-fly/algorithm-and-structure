package tree

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only root node",
			args: args{t: NewTreeFromArray(level1Arr)},
			want: true,
		},
		{
			name: "2-level non-symmetric",
			args: args{t: NewTreeFromArray(level2Arr)},
			want: false,
		},
		{
			name: "2-level symmetric",
			args: args{t: NewTreeFromArray([]any{1, 2, 2})},
			want: true,
		},
		{
			name: "4-level symmetric",
			args: args{t: NewTreeFromArray(level4SymmetricArr)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetric(tt.args.t.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSymmetricByBFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only root node",
			args: args{t: NewTreeFromArray(level1Arr)},
			want: true,
		},
		{
			name: "2-level non-symmetric",
			args: args{t: NewTreeFromArray(level2Arr)},
			want: false,
		},
		{
			name: "2-level symmetric",
			args: args{t: NewTreeFromArray([]any{1, 2, 2})},
			want: true,
		},
		{
			name: "4-level symmetric",
			args: args{t: NewTreeFromArray(level4SymmetricArr)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetricByBFS(tt.args.t.root); got != tt.want {
				t.Errorf("IsSymmetricByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
