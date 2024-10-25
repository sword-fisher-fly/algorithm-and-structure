package graph

import "testing"

func TestSequenceReconstruction(t *testing.T) {
	type args struct {
		n         int
		sequences [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{n: 4, sequences: [][]int{{1, 2}, {2, 3}, {3, 4}}},
			want: true,
		},
		{
			name: "case 2",
			args: args{n: 3, sequences: [][]int{{1, 2}}},
			want: false,
		},
		{
			name: "case 3",
			args: args{n: 3, sequences: [][]int{{1, 2}, {1, 3}, {2, 3}}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SequenceReconstruction(tt.args.n, tt.args.sequences); got != tt.want {
				t.Errorf("SequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
