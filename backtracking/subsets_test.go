package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1 item",
			args: args{arr: []int{1}},
			want: [][]int{{}, {1}},
		},

		{
			name: "2 items",
			args: args{arr: []int{1, 2}},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "3 items",
			args: args{arr: []int{1, 2, 3}},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subsets(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsetsWithDup(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "2 items",
			args: args{arr: []int{1, 1}},
			want: [][]int{{}, {1}, {1, 1}},
		},
		{
			name: "3 items",
			args: args{arr: []int{1, 2, 2}},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetsWithDup(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsetsWithDupII(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "2 items",
			args: args{arr: []int{1, 1}},
			want: [][]int{{}, {1}, {1, 1}},
		},
		{
			name: "3 items",
			args: args{arr: []int{1, 2, 2}},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetsWithDupII(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubsetsWithDupII() = %v, want %v", got, tt.want)
			}
		})
	}
}
