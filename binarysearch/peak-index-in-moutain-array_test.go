package binarysearch

import (
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				arr: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				arr: []int{0, 10, 5, 2},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				arr: []int{1, 3, 5, 4, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PeakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("PeakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeakIndexInMountainArrayII(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				arr: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				arr: []int{0, 10, 5, 2},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				arr: []int{1, 3, 5, 4, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PeakIndexInMountainArrayII(tt.args.arr); got != tt.want {
				t.Errorf("PeakIndexInMountainArrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}
