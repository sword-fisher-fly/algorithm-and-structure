package math

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 4, 3},
			},
			want: []int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
