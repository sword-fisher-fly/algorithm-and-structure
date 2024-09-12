package binarysearch

import "testing"

// no duplicated key in the rotated array
func TestBinarySearchInRotatedArray(t *testing.T) {
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
			name: "hit an item in rotated array",
			args: args{
				arr:    []int{4, 5, 6, 7, 0, 1, 2, 3},
				target: 7,
			},
			want: 3,
		},
		{
			name: "empty array",
			args: args{
				arr:    []int{},
				target: 7,
			},
			want: -1,
		},
		{
			name: "target does not appear in the rotated array",
			args: args{
				arr:    []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchInRotatedArray(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearchInRotatedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
