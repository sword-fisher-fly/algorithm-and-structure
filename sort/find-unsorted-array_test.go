package sort

import (
	"testing"
)

func TestFindUnsortedSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{nums: []int{2, 6, 4, 8, 10, 9, 15}},
			want: 5,
		},
		{
			name: "case 2",
			args: args{nums: []int{1, 3, 2, 2, 2, 5}},
			want: 4,
		},
		{
			name: "case 3",
			args: args{nums: []int{1, 4, 2, 2, 4, 5, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUnsortedSubarray(tt.args.nums); got != tt.want {
				t.Errorf("FindUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUnSortedSubarrayII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{nums: []int{2, 6, 4, 8, 10, 9, 15}},
			want: 5,
		},
		{
			name: "case 2",
			args: args{nums: []int{1, 3, 2, 2, 2, 5}},
			want: 4,
		},
		{
			name: "case 3",
			args: args{nums: []int{1, 4, 2, 2, 4, 5, 4}},
			want: 6,
		},
		{
			name: "case 4",
			args: args{nums: []int{1, 2, 3}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUnSortedSubarrayII(tt.args.nums); got != tt.want {
				t.Errorf("FindUnSortedSubarrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}
