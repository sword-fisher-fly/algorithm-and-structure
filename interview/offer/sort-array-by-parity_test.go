package offer

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort array by parity: [4, 2, 5, 7]",
			args: args{
				nums: []int{4, 2, 5, 7},
			},
			want: []int{4, 5, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort array by parity: [4, 2, 5, 7]",
			args: args{
				nums: []int{4, 2, 5, 7},
			},
			want: []int{4, 5, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayByParityII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
