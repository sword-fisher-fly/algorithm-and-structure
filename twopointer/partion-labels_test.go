package twopointer

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ababcbacadefegdehijhklij",
			args: args{s: "ababcbacadefegdehijhklij"},
			want: []int{9, 7, 8},
		},
		{
			name: "abcdefghijklmnopqrstuvwxyz",
			args: args{s: "abcdefghijklmnopqrstuvwxyz"},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartitionLabels(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
