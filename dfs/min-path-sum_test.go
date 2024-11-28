package dfs

import (
	"testing"
)

// func TestMinPathSum(t *testing.T) {
// 	type args struct {
// 		grid [][]int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "case 1",
// 			args: args{
// 				grid: [][]int{
// 					{1, 3, 1},
// 					{1, 5, 1},
// 					{4, 2, 1},
// 				},
// 			},
// 			want: 7,
// 		},
// 		{
// 			name: "case 2",
// 			args: args{
// 				grid: [][]int{
// 					{1, 4, 8, 6, 2, 2, 1, 7},
// 					{4, 7, 3, 1, 4, 5, 5, 1},
// 					{8, 8, 2, 1, 1, 8, 0, 1},
// 					{8, 9, 2, 9, 8, 0, 8, 9},
// 					{5, 7, 5, 7, 1, 8, 5, 5},
// 					{7, 0, 9, 4, 5, 6, 5, 6},
// 					{4, 9, 9, 7, 9, 1, 9, 0},
// 				},
// 			},
// 			want: 0,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := MinPathSum(tt.args.grid); got != tt.want {
// 				t.Errorf("MinPathSum() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestMinPathSumByRecursive(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				grid: [][]int{
					{1, 4, 8, 6, 2, 2, 1, 7},
					{4, 7, 3, 1, 4, 5, 5, 1},
					{8, 8, 2, 1, 1, 8, 0, 1},
					{8, 9, 2, 9, 8, 0, 8, 9},
					{5, 7, 5, 7, 1, 8, 5, 5},
					{7, 0, 9, 4, 5, 6, 5, 6},
					{4, 9, 9, 7, 9, 1, 9, 0},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPathSumByRecursive(tt.args.grid); got != tt.want {
				t.Errorf("MinPathSumByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
