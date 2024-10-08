package interview

import "testing"

func TestFindMedianInTwoSortedArrays(t *testing.T) {
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
			name: "total number is odd",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "total number is even",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMedianInTwoSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("FindMedianInTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
