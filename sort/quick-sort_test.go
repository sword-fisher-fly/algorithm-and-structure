package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "5 items",
			args: args{arr: []int{3, 2, 1, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "1 items",
			args: args{arr: []int{1}},
			want: []int{1},
		},
		{
			name: "2 items",
			args: args{arr: []int{2, 1}},
			want: []int{1, 2},
		},
		{
			name: "10 items",
			args: args{arr: []int{3, 2, 1, 4, 5, 10, 8, 7, 6, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortII(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "5 items",
			args: args{arr: []int{3, 2, 1, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "1 items",
			args: args{arr: []int{1}},
			want: []int{1},
		},
		{
			name: "2 items",
			args: args{arr: []int{2, 1}},
			want: []int{1, 2},
		},
		{
			name: "10 items",
			args: args{arr: []int{3, 2, 1, 4, 5, 10, 8, 7, 6, 9}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSortII(tt.args.arr)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("QuickSortII() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}
