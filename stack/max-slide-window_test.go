package stack

import (
	"reflect"
	"testing"
)

func TestMaxInSlideWindow(t *testing.T) {
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
			name: "standard example",
			args: args{arr: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "20 items",
			args: args{arr: []int{3, -1, 2, 3, 4, -1, 9, 6, 7, 18, 4, 5, 7, -3, 20, 12, 8, 5, -1, 16}, k: 5},
			// 1) 3, -1,2,3,4
			// 2) -1,2,3,4,-1
			// 3) 2,3,4,-1,9
			// 4) 3,4,-1,9,6
			// 5) 4,-1,9,6,7
			// 6) -1,9,6,7,18
			// 7) 9,6,7,18,4
			// 8) 6,7,18,4,5
			// 9) 7,18,4,5,7
			// 10) 18,4,5,7,-3
			// 11) 4,5,7,-3,20
			// 12) 5, 7, -3, 20, 12
			// 13) 7, -3, 20, 12, 8
			// 14) -3, 20, 12, 8, 5
			// 15) 20, 12, 8, 5, -1
			// 16) 12, 8, 5, -1, 16
			want: []int{4, 4, 9, 9, 9, 18, 18, 18, 18, 18, 20, 20, 20, 20, 20, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInSlideWindow(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxInSlideWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
