package offer

import (
	"testing"
)

func TestFindTargetSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindTargetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTargetSumByDP(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSumByDP(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindTargetSumByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
