package backtracking

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "extract 2 items from 1..5",
			args: args{
				n: 5,
				k: 2,
			},
			want: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{1, 5},
				{2, 3},
				{2, 4},
				{2, 5},
				{3, 4},
				{3, 5},
				{4, 5},
			},
		},
		{
			name: "extract 3 items from 1..5",
			args: args{
				n: 5,
				k: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 2, 5},
				{1, 3, 4},
				{1, 3, 5},
				{1, 4, 5},
				{2, 3, 4},
				{2, 3, 5},
				{2, 4, 5},
				{3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combine() = %v, want %v", got, tt.want)
			}

			if got := CombineII(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombineII() = %v, want %v", got, tt.want)
			}
		})
	}
}
