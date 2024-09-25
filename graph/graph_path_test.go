package graph

import (
	"reflect"
	"testing"
)

func TestFindAllGraphPath(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				graph: [][]int{
					{1, 2},
					{3},
					{3},
					{},
				},
			},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			name: "case 2",
			args: args{
				graph: [][]int{
					{4, 3, 1},
					{3, 2, 4},
					{3},
					{4},
					{},
				},
			},
			want: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			name: "case 3 no path",
			args: args{
				graph: [][]int{
					{1, 2},
					{},
					{},
					{},
				},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllGraphPath(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllGraphPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
