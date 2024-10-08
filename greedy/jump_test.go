package greedy

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can jump",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "can not jump",
			args: args{
				nums: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.args.nums); got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJumpMinSteps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "can jump min step is 2",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "can jump min step is 3",
			args: args{
				nums: []int{2, 2, 1, 1, 4},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpMinSteps(tt.args.nums); got != tt.want {
				t.Errorf("JumpMinSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
