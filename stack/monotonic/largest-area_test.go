package monotonic

import (
	"testing"
)

func TestLargestRectangleAreaByForce(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 item",
			args: args{heights: []int{1}},
			want: 1,
		},
		{
			name: "2 items",
			args: args{heights: []int{1, 2}},
			want: 2,
		},
		{
			name: "2 1 5 6 2 3",
			args: args{heights: []int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleAreaByForce(tt.args.heights); got != tt.want {
				t.Errorf("LargestRectangleAreaByForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestRectangleAreaByTwoPointer(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 item",
			args: args{heights: []int{1}},
			want: 1,
		},
		{
			name: "2 items",
			args: args{heights: []int{1, 2}},
			want: 2,
		},
		{
			name: "2 1 5 6 2 3",
			args: args{heights: []int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleAreaByTwoPointer(tt.args.heights); got != tt.want {
				t.Errorf("LargestRectangleAreaByTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 item",
			args: args{heights: []int{1}},
			want: 1,
		},
		{
			name: "2 items",
			args: args{heights: []int{1, 2}},
			want: 2,
		},
		{
			name: "2 1 5 6 2 3",
			args: args{heights: []int{2, 1, 5, 6, 2, 3}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
