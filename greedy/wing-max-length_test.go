package greedy

import (
	"testing"
)

func TestWiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only two elements in the array",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "the array with the two same elements",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "the array in ascending order",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 2,
		},
		{
			name: "the array in descending order",
			args: args{
				nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 2,
		},
		{
			name: "the array with the same adjacent elements",
			args: args{
				nums: []int{1, 1, 1, 2, 1, 1, 3, 1, 1},
			},
			want: 5,
		},
		{
			name: "random array",
			args: args{
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("WiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWiggleMaxLengthByDP(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only two elements in the array",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "the array with the two same elements",
			args: args{nums: []int{1, 1}},
			want: 1,
		},
		{
			name: "the array in ascending order",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: 2,
		},
		{
			name: "the array in descending order",
			args: args{
				nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 2,
		},
		{
			name: "the array with the same adjacent elements",
			args: args{
				nums: []int{1, 1, 1, 2, 1, 1, 3, 1, 1},
			},
			want: 5,
		},
		{
			name: "random array",
			args: args{
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WiggleMaxLengthByDP(tt.args.nums); got != tt.want {
				t.Errorf("WiggleMaxLengthByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
