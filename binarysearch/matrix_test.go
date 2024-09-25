package binarysearch

import (
	"reflect"
	"testing"
)

func TestFindTargetInSortedMatrixII(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "empty matrix",
			args: args{matrix: [][]int{}, target: 1},
			want: [2]int{-1, -1},
		},
		{
			name: "target exists in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 8,
			},
			want: [2]int{1, 2},
		},
		{
			name: "target exists in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 11,
			},
			want: [2]int{0, 3},
		},
		{
			name: "target exists in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 4,
			},
			want: [2]int{0, 1},
		},
		{
			name: "target exists in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 12,
			},
			want: [2]int{1, 3},
		},
		{
			name: "target does not exist in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			want: [2]int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetInSortedMatrixII(tt.args.matrix, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTargetInSortedMatrixII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTargetInSortedMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "empty matrix",
			args: args{matrix: [][]int{}, target: 1},
			want: [2]int{-1, -1},
		},
		{
			name: "target exists in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{22, 25, 28, 30, 31},
					{34, 36, 37, 40, 41},
					{45, 47, 49, 50, 53},
					{56, 58, 60, 62, 63},
				},
				target: 28,
			},
			want: [2]int{1, 2},
		},
		{
			name: "target doesn't exist in the normal matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{22, 25, 28, 30, 31},
					{34, 36, 37, 40, 41},
					{45, 47, 49, 50, 53},
					{56, 58, 60, 62, 63},
				},
				target: 16,
			},
			want: [2]int{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetInSortedMatrix(tt.args.matrix, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTargetInSortedMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
