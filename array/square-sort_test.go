package array

import (
	"reflect"
	"testing"
)

func TestSquareSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "only positive items",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "only negative items",
			args: args{
				nums: []int{-1, -2, -3, -4, -5},
			},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "both negative and positive items",
			args: args{
				nums: []int{-5, -3, 1, 2, 6, 7},
			},
			want: []int{1, 4, 9, 25, 36, 49},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
