package matrix

import (
	"reflect"
	"testing"
)

func TestRotateByClockwise(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3*3 matrix",
			args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateByClockwise(tt.args.matrix)

			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("RotateByClockwise() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}

func TestRotateByAntiClockwise(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3*3 matrix",
			args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateByAntiClockwise(tt.args.matrix)

			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("RotateByAntiClockwise() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}

func TestRotateBy180Clockwise(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3*3 matrix",
			args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateBy180Clockwise(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("RotateBy180Clockwise() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}

func TestRotateBy180ClockwiseII(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3*3 matrix",
			args: args{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateBy180ClockwiseII(tt.args.matrix)

			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("RotateBy180ClockwiseII() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
