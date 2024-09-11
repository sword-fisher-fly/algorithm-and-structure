package array

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "3*3 matrix",
			args: args{
				n: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
