package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hit an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 3,
			},
			want: 2,
		},
		{
			name: "miss an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchInSortedArray(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchInSortedArrayII(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hit an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 3,
			},
			want: 2,
		},
		{
			name: "missing an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchInSortedArrayII(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchInSortedArrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchInSortedArrayRecursive(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hit an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 3,
			},
			want: 2,
		},
		{
			name: "miss an item",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
				target: 6,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchInSortedArrayRecursive(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchInSortedArrayRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
