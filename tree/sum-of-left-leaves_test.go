package tree

import (
	"testing"
)

func TestSumOfLeftLeavesByRecursive(t *testing.T) {
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
			args: args{t: NewTreeFromArray([]any{2, 1, 3})},
			want: 1,
		},
		{
			name: "random tree",
			args: args{t: NewTreeFromArray([]any{2, 9, 20, 5, 7, 15, 1})},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfLeftLeavesByRecursive(tt.args.t.root); got != tt.want {
				t.Errorf("SumOfLeftLeavesByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfLeftLeavesByIteration(t *testing.T) {
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
			args: args{t: NewTreeFromArray([]any{2, 1, 3})},
			want: 1,
		},
		{
			name: "random tree",
			args: args{t: NewTreeFromArray([]any{2, 9, 20, 5, 7, 15, 1})},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfLeftLeavesByIteration(tt.args.t.root); got != tt.want {
				t.Errorf("SumOfLeftLeavesByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
