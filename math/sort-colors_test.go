package math

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	type args struct {
		colors []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1,0,2",
			args: args{colors: []int{1, 0, 2}},
			want: []int{0, 1, 2},
		},
		{
			name: "2,0,2,1,1,0",
			args: args{colors: []int{2, 0, 2, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortColors(tt.args.colors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
