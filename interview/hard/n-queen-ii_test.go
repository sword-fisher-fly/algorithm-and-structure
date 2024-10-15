package hard

import "testing"

func TestTotalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4 * 4 board has 2 solutions",
			args: args{
				n: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalNQueens(tt.args.n); got != tt.want {
				t.Errorf("TotalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
