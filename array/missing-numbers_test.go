package array

import (
	"reflect"
	"testing"
)

func TestFindMissingNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "5 items",
			args: args{nums: []int{3, 4, 3, 1, 5}},
			want: []int{2},
		},
		{
			name: "8 items",
			args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			want: []int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMissingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
