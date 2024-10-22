package offer

import (
	"reflect"
	"testing"
)

func TestCountBitsIII(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "n=2",
			args: args{n: 2},
			want: []int{1, 1},
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: []int{1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBitsIII(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBitsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
