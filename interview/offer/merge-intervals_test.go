package offer

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},

		{
			name: "case 2",
			args: args{
				intervals: [][]int{{1, 4}, {2, 3}},
			},
			want: [][]int{{1, 4}},
		},

		{
			name: "case 3",
			args: args{
				intervals: [][]int{{1, 4}, {0, 4}},
			},
			want: [][]int{{0, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
