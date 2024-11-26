package huawei

import (
	"reflect"
	"testing"
)

func TestMergeRecord(t *testing.T) {
	type args struct {
		pairs [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "case 1",
			args: args{
				pairs: [][2]int{
					{0, 1},
					{0, 2},
					{1, 3},
					{2, 4},
				},
			},
			want: [][2]int{
				{0, 3},
				{1, 3},
				{2, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeRecord(tt.args.pairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
