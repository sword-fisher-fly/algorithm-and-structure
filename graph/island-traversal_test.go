package graph

import (
	"testing"
)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIslandsByBFS(tt.args.grid); got != tt.want {
				t.Errorf("FindIslandsByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIslandsByDFS(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIslandsByDFS(tt.args.grid); got != tt.want {
				t.Errorf("FindIslandsByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxIslandArea(t *testing.T) {
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
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIslandArea(tt.args.grid); got != tt.want {
				t.Errorf("MaxIslandArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxIsolatedIslandAreaByDFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one isolated islands in the 5*4 grid",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "two isolated islands in the 7*8 grid",
			args: args{
				grid: [][]int{
					{0, 1, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0, 1, 1},
					{0, 1, 1, 1, 0, 1, 1, 1},
					{0, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 1, 1, 0},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIsolatedIslandAreaByDFS(tt.args.grid); got != tt.want {
				t.Errorf("MaxIsolatedIslandAreaByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxIsolatedIslandAreaByBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one isolated islands in the 5*4 grid",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 1},
					{0, 0, 0, 1, 1},
				},
			},
			want: 1,
		},
		{
			name: "two isolated islands in the 7*8 grid",
			args: args{
				grid: [][]int{
					{0, 1, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0, 1, 1},
					{0, 1, 1, 1, 0, 1, 1, 1},
					{0, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 1, 1, 0},
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIsolatedIslandAreaByBFS(tt.args.grid); got != tt.want {
				t.Errorf("MaxIsolatedIslandAreaByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLargestAreaIslandByBFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2*2 grid",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: 3,
		},
		{
			name: "full of all land in the 5*4 grid",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
				},
			},
			want: 20,
		},
		{
			name: "5*4 grid with one flip which can connect the two islands",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 1, 1, 1},
				},
			},
			want: 9,
		},
		{
			name: "two isolated islands in the 7*8 grid",
			args: args{
				grid: [][]int{
					{0, 1, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0, 1, 1},
					{0, 1, 1, 1, 0, 1, 1, 1},
					{0, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 1, 1, 0},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargestAreaIslandByBFS(tt.args.grid); got != tt.want {
				t.Errorf("FindLargestAreaIslandByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLargestAreaIslandByDFS(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2*2 grid",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: 3,
		},
		{
			name: "full of all land in the 5*4 grid",
			args: args{
				grid: [][]int{
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1},
				},
			},
			want: 20,
		},
		{
			name: "5*4 grid with one flip which can connect the two islands",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0, 0},
					{1, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 1, 1, 1},
				},
			},
			want: 9,
		},
		{
			name: "two isolated islands in the 7*8 grid",
			args: args{
				grid: [][]int{
					{0, 1, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0, 1, 1},
					{0, 1, 1, 1, 0, 1, 1, 1},
					{0, 0, 0, 0, 1, 0, 0, 0},
					{0, 1, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 1, 1, 0},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLargestAreaIslandByDFS(tt.args.grid); got != tt.want {
				t.Errorf("FindLargestAreaIslandByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxIslandAreaII(t *testing.T) {
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
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxIslandAreaII(tt.args.grid); got != tt.want {
				t.Errorf("MaxIslandAreaII() = %v, want %v", got, tt.want)
			}
		})
	}
}
