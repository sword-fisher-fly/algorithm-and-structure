package stack

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	type args struct {
		arr []*Interval
	}
	tests := []struct {
		name string
		args args
		want []*Interval
	}{
		{
			name: "case 1",
			args: args{
				arr: []*Interval{{1, 4}, {2, 3}},
			},
			want: []*Interval{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
