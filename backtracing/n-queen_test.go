package backtracking

import (
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1*1 board",
			args: args{n: 1},
			want: [][]string{{"Q"}},
		},
		{
			name: "4*4 board",
			args: args{n: 4},
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
