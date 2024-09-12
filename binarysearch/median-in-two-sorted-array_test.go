package binarysearch

import (
	"reflect"
	"testing"
)

func Test_mergeTwoSortedArray(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "an empty array and a non-empty array",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "total 15 items in nums1 and nums2",
			args: args{
				nums1: []int{1, 3, 5, 7, 9, 11, 12, 13, 14, 15},
				nums2: []int{2, 4, 6, 8, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoSortedArray(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMedianInTwoSortedArrayByForce(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "an empty array and a non-empty array",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "total number of the two arrays is odd",
			args: args{
				nums1: []int{1, 3, 5, 7, 9, 11, 12, 13, 14, 15},
				nums2: []int{2, 4, 6, 8, 10},
			},
			want: 8,
		},
		{
			name: "total number of the two arrays is even",
			args: args{
				nums1: []int{1, 3, 5, 7, 9, 11, 12, 13, 14},
				nums2: []int{2, 4, 6, 8, 10},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianInTwoSortedArrayByForce(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianInTwoSortedArrayByForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
