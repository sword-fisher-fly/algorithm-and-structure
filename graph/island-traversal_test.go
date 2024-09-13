package graph

import "testing"

func TestFindIslandsByBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3 islands in the 5*4 grid",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				},
			},
			want: 3,
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIslandsByBFS(tt.args.grid); got != tt.want {
				t.Errorf("FindIslandsByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
