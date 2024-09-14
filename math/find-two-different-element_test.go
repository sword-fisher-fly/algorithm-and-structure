package math

import (
	"reflect"
	"testing"
)

func TestFindTwoDifferentElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		want   [2]int
		wantII [2]int
	}{
		{
			name:   "10 elements in a array",
			args:   args{nums: []int{2, 3, 6, 5, 19, 5, 3, 2}},
			want:   [2]int{6, 19},
			wantII: [2]int{19, 6},
		},
		{
			name:   "20 elements in a array",
			args:   args{nums: []int{2, 3, 6, 5, 19, 5, 3, 2, 11, 116, 17, 18, 30, 17, 11, 116, 30, 18}},
			want:   [2]int{6, 19},
			wantII: [2]int{19, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindTwoDifferentElement(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.wantII) {
				t.Errorf("FindTwoDifferentElement() got = %v, want %v or %v", got, tt.want, tt.wantII)
			}
		})
	}
}
