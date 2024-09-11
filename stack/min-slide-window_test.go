package stack

import (
	"reflect"
	"testing"
)

func TestMinInSlideWindow(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{4, 3, 5, 4, 3, 3, 6, 7},
				k:   3,
			},
			want: []int{3, 3, 3, 3, 3, 3},
		},
		{
			name: "standard example",
			args: args{arr: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: []int{-1, -3, -3, -3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInSlideWindow(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinInSlideWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
