package backtracking

import (
	"reflect"
	"testing"
)

func Test_CombineSumWithRepeat(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "case2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := CombineSumWithRepeat(tt.args.candidates, tt.args.target)

			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("CombineSumWithRepeat(%v,%v): got %v, want %v", tt.args.candidates, tt.args.target, ret, tt.want)
			}
		})
	}
}

func Test_CombineSumWithRepeatII(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "case2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := CombineSumWithRepeatII(tt.args.candidates, tt.args.target)

			if !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("CombineSumWithRepeatII(%v,%v): got %v, want %v", tt.args.candidates, tt.args.target, ret, tt.want)
			}
		})
	}
}

func TestCombineSumWithRepeatOptimized(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "case2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineSumWithRepeatOptimized(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombineSumWithRepeatOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}
