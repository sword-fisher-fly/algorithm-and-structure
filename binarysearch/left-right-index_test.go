package binarysearch

import (
	"testing"
)

func TestSearchLeftRightIndex(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "nums=[1,2,3,3,3,4,5], target=3",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: 4,
		},

		{
			name: "nums=[1,2,3,3,3,4,5], target=6",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want:  -1,
			want1: -1,
		},

		{
			name: "nums=[1,2,3,3,3,4,5], target=1",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 1,
			},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SearchLeftRightIndex(tt.args.nums, tt.args.target)
			if got != tt.want {
				t.Errorf("SearchLeftRightIndex() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SearchLeftRightIndex() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSearchLeftRightIndexII(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "nums=[1,2,3,3,3,4,5], target=3",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 3,
			},
			want:  2,
			want1: 4,
		},

		{
			name: "nums=[1,2,3,3,3,4,5], target=6",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 6,
			},
			want:  -1,
			want1: -1,
		},

		{
			name: "nums=[1,2,3,3,3,4,5], target=1",
			args: args{
				nums:   []int{1, 2, 3, 3, 3, 4, 5},
				target: 1,
			},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SearchLeftRightIndexII(tt.args.nums, tt.args.target)
			if got != tt.want {
				t.Errorf("SearchLeftRightIndexII() got = %v, want %v", got, tt.want)
			}

			if got1 != tt.want1 {
				t.Errorf("SearchLeftRightIndexII() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
