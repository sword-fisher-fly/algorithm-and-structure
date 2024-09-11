package array

import (
	"reflect"
	"testing"
)

// 1) Input: [[1,3],[2,6],[8,10],[15,18]]
//    Output: [[1,6],[8,10],[15,18]]
// 2) Input: [[1,4],[4,5]]
//    Output: [1,5]

func TestMergeIntervals(t *testing.T) {
	type args struct {
		intervals [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "2 intervals",
			args: args{
				intervals: [][2]int{{1, 3}, {2, 4}},
			},
			want: [][2]int{{1, 4}},
		},
		{
			name: "3 intervals",
			args: args{
				intervals: [][2]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][2]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "4 intervals",
			//[1,3],[2,6],[8,10],[15,18]
			args: args{
				intervals: [][2]int{{1, 3}, {2, 4}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][2]int{{1, 6}, {8, 10}, {15, 18}},
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
