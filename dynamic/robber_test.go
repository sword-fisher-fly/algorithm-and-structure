package dynamic

import (
	"example/alg-structure/tree"
	"testing"
)

func TestRobberHouseOnLine(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only one house",
			args: args{arr: []int{1}},
			want: 1,
		},
		{
			name: "two houses",
			args: args{arr: []int{1, 2}},
			want: 2,
		},
		{
			name: "three houses",
			args: args{arr: []int{1, 2, 3}},
			want: 4,
		},
		{
			name: "ten house",
			// 1,3,5,7,10
			// 2,4,6,8,10
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobberHouseOnLine(tt.args.arr); got != tt.want {
				t.Errorf("RobberHouseOnLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobberHoustOnLoop(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only one house",
			args: args{arr: []int{1}},
			want: 1,
		},
		{
			name: "two houses",
			args: args{arr: []int{1, 2}},
			want: 2,
		},
		{
			name: "three houses",
			args: args{arr: []int{1, 2, 3}},
			want: 3,
		},
		{
			name: "ten house",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			// 1,3,5,6,9
			// 2,4,6,8,10
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobberHoustOnLoop(tt.args.arr); got != tt.want {
				t.Errorf("RobberHoustOnLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRobberHouseOnTree(t *testing.T) {
	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// [3,2,3,-1,3,-1,1] -> 7
		// [3,4,5,1,3,-1,1] -> 9
		{
			name: "case 1",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 2, 3, -1, 3, -1, 1}),
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 4, 5, 1, 3, -1, 1}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobberHouseOnTree(tt.args.t.Root()); got != tt.want {
				t.Errorf("RobberHouseOnTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RobberHouseOnTreeWithMemory(t *testing.T) {
	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 2, 3, -1, 3, -1, 1}),
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 4, 5, 1, 3, -1, 1}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobberHouseOnTreeWithMemory(tt.args.t.Root()); got != tt.want {
				t.Errorf("RobberHouseOnTreeWithMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RobberHouseOnTreeByDP(t *testing.T) {
	type args struct {
		t *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 2, 3, -1, 3, -1, 1}),
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				t: tree.NewTreeFromArray([]any{3, 4, 5, 1, 3, -1, 1}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobberHouseOnTreeByDP(tt.args.t.Root()); got != tt.want {
				t.Errorf("RobberHouseOnTreeByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
