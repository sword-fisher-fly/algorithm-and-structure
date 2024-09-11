package twopointer

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][4]int
	}{
		{
			name: "standard array",
			args: args{nums: []int{1, 0, -1, 0, -2, 2}},
			want: [][4]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FourSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
