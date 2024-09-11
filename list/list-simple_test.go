package list

import (
	"reflect"
	"testing"
)

func TestToArray(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArray(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "9 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "9 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseListRecursive(tt.args.head); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("ReverseListRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseListRecursiveII(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "9 items list to array",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 6, 7, 8, 9, 10}),
			},
			want: []int{10, 9, 8, 7, 6, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseListRecursiveII(tt.args.head); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("ReverseListRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwapPairs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9},
		},
		{
			name: "11 items",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			},
			want: []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SwapPairs(NewListFromArray(tt.args.arr)); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("SwapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
