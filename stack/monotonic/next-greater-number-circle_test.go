package monotonic

import (
	"reflect"
	"testing"
)

func TestNextGreaterNumberInCircleArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1 2 1",
			args: args{nums: []int{1, 2, 1}},
			want: []int{2, -1, 2},
		},
		{
			name: "1 2 3 4 3 1",
			args: args{nums: []int{1, 2, 3, 4, 3, 1}},
			want: []int{2, 3, 4, -1, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterNumberInCircleArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterNumberInCircleArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNextGreaterNumberInCircleArrayII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1 2 1",
			args: args{nums: []int{1, 2, 1}},
			want: []int{2, -1, 2},
		},
		{
			name: "1 2 3 4 3 1",
			args: args{nums: []int{1, 2, 3, 4, 3, 1}},
			want: []int{2, 3, 4, -1, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextGreaterNumberInCircleArrayII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextGreaterNumberInCircleArrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}
