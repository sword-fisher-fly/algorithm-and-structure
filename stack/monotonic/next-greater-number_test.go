package monotonic

import (
	"reflect"
	"testing"
)

func TestNextGreaterNumber(t *testing.T) {
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
			name: "standard array",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			name: "descend array",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{4, 3, 2, 1},
			},
			want: []int{-1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterNumber(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextGreaterElementII(t *testing.T) {
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
			name: "standard array",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			name: "descend array",
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{4, 3, 2, 1},
			},
			want: []int{-1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterElementII(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}
