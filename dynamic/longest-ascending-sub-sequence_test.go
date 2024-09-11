package dynamic

import (
	"testing"
)

func TestLongestAscendingSubSequence(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 item",
			args: args{arr: []int{1}},
			want: 1,
		},
		{
			name: "2 items",
			args: args{arr: []int{1, 2}},
			want: 2,
		},
		{
			name: "6 - 4 items",
			args: args{arr: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "",
			args: args{arr: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			// [2,3,7,101]
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestAscendingSubSequence(tt.args.arr); got != tt.want {
				t.Errorf("LongestAscendingSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestContinuousAscendingSubSequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 continuous items",
			args: args{nums: []int{1, 2}},
			want: 2,
		},
		{
			name: "3 continuous items of 5 items",
			args: args{nums: []int{1, 3, 5, 4, 7}},
			want: 3,
		},
		{
			name: "9 continuous items of 10 items",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestContinuousAscendingSubSequence(tt.args.nums); got != tt.want {
				t.Errorf("LongestContinuousAscendingSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
