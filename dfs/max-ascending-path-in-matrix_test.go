package dfs

import "testing"

func TestMaxPathAscendingInMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxPathAscendingInMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("MaxPathAscendingInMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
