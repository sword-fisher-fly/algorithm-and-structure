package backtracking

import (
	"reflect"
	"testing"
)

func TestFindAscendingSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sorted array",
			args: args{
				nums: []int{4, 6, 7, 7},
			},
			want: [][]int{{4, 6}, {4, 6, 7}, {4, 7}, {6, 7}},
		},
		{
			name: "non-sorted array",
			args: args{
				nums: []int{4, 7, 6, 7},
			},
			want: [][]int{{4, 7}, {4, 6}, {4, 6, 7}, {6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAscendingSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAscendingSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
