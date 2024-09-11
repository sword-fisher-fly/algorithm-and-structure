package tree

import "testing"

func TestIsBalancedTreeByRecursive(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-level tree",
			args: args{t: level2FullTree},
			want: true,
		},
		{
			name: "4-level tree",
			args: args{t: level4FullTree},
			want: true,
		},
		{
			name: "random tree",
			args: args{t: level4RandomTree},
			want: false,
		},
		{
			name: "symmetric tree",
			args: args{t: level4SymmetricTree},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalancedTreeByRecursive(tt.args.t.root); got != tt.want {
				t.Errorf("IsBalancedTreeByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
