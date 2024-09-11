package sort

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		want   []int
		wantII []int
	}{
		{
			name: "no same frequency items",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},

			want: []int{1, 2},
		},
		{
			name: "same frequency items",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6},
				k:    3,
			},
			want:   []int{1, 2, 3},
			wantII: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.wantII) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}

		})
	}
}
