package backtracking

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "2 items",
			args: args{s: "ab"},
			want: []string{"ab", "ba"},
		},
		{
			name: "3 items",
			args: args{s: "abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutationWithDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "two same items",
			args: args{nums: []int{1, 1}},
			want: [][]int{{1, 1}},
		},
		{
			name: "three same items",
			args: args{nums: []int{1, 1, 2}},
			want: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermutationWithDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermutationWithDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
