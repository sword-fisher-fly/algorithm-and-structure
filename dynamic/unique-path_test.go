package dynamic

import (
	"testing"
)

func TestUniquePath(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1*1",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "2*2",
			args: args{
				m: 2,
				n: 2,
			},
			want: 2,
		},
		{
			name: "3*3",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "7*3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePath(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("UniquePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniquePathII(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1*1",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "2*2",
			args: args{
				m: 2,
				n: 2,
			},
			want: 2,
		},
		{
			name: "3*3",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			name: "7*3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathII(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("UniquePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniquePathWithObstacles(t *testing.T) {
	type args struct {
		m         int
		n         int
		obstacles [][]bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3*3 with 1 obstacle",
			args: args{
				m: 3,
				n: 3,
				obstacles: [][]bool{
					{false, false, false},
					{false, true, false},
					{false, false, false},
				},
			},
			want: 2,
		},
		{
			name: "3*4 with 1 obstacle",
			args: args{
				m: 3,
				n: 4,
				obstacles: [][]bool{
					{false, false, false, false},
					{false, true, false, false},
					{false, false, false, false},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathWithObstacles(tt.args.m, tt.args.n, tt.args.obstacles); got != tt.want {
				t.Errorf("UniquePathWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
