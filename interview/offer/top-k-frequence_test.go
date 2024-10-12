package offer

import (
	"reflect"
	"testing"
)

func TestTopKFrequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2,2,3,1,1",
			args: args{
				nums: []int{2, 2, 3, 1, 1},
				k:    2,
			},
			want: []int{1, 2},
		},

		{
			name: "2,2,3,1,1,4,5,9,8,4,5,4",
			args: args{
				nums: []int{2, 2, 3, 1, 1, 4, 5, 9, 8, 4, 5, 4},
				k:    2,
			},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequence(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopKFrequenceII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1,1,2,2,2,3,3,3,3,4,5,6,7",
			args: args{
				nums: []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequenceII(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequenceII() = %v, want %v", got, tt.want)
			}
		})
	}
}
