package hash

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int64
		target int64
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "sorted array",
			args: args{
				nums:   []int64{2, 7, 11, 15},
				target: 9,
			},
			want: [2]int{0, 1},
		},
		{
			name: "unsorted array",
			args: args{
				nums:   []int64{3, 2, 4},
				target: 6,
			},
			want: [2]int{1, 2},
		},
		{
			name: "sorted array with two same item",
			args: args{
				nums:   []int64{1, 2, 3, 3, 7},
				target: 6,
			},
			want: [2]int{2, 3},
		},
		{
			name: "unsorted array with two same item",
			args: args{
				nums:   []int64{3, 3, 1, 2, 7},
				target: 6,
			},
			want: [2]int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
