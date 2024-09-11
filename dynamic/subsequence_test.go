package dynamic

import (
	"testing"
)

func TestMaxLengthOfIncreasingSubSequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5 items",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "10 items",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 10,
		},
		{
			name: "mix array",
			args: args{
				nums: []int{3, 2, 4, 7, 8, 6, 8, 9, 10},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxLengthOfIncreasingSubSequence(tt.args.nums); got != tt.want {
				t.Errorf("MaxLengthOfIncreasingSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxRepeatedSubArray(t *testing.T) {
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
			name: "5 items",
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			want: 3,
		},
		{
			name: "10 items",
			args: args{
				nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				nums2: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxRepeatedSubArray(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("MaxRepeatedSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubSequence(t *testing.T) {
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
			name: "5 items",
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			want: 3,
		},
		{
			name: "10 items",
			args: args{
				nums1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				nums2: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubSequence(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("LongestCommonSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubSequenceII(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the same string",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: 3,
		},
		{
			name: "different string 1",
			args: args{
				s1: "abc",
				s2: "abd",
			},
			want: 2,
		},
		{
			name: "different string 2",
			args: args{
				s1: "abcde",
				s2: "ace",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubSequenceII(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LongestCommonSubSequenceII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSumOfSubSequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumOfSubSequence(tt.args.nums); got != tt.want {
				t.Errorf("MaxSumOfSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "the same string",
		// 	args: args{
		// 		s: "abc",
		// 		t: "abc",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "different string 1",
		// 	args: args{
		// 		s: "abc",
		// 		t: "abd",
		// 	},
		// 	want: false,
		// },
		{
			name: "different string 2",
			args: args{
				s: "ace",
				t: "abcde",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
