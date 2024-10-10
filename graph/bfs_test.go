package graph

import (
	"strings"
	"testing"
)

func prettyPrintGrid(visitedPath [][]bool) string {
	var sb strings.Builder
	for i := 0; i < len(visitedPath); i++ {
		for j := 0; j < len(visitedPath[0]); j++ {
			if visitedPath[i][j] {
				sb.WriteByte('1')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func Test_bfs(t *testing.T) {
	type args struct {
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}
	tests := []struct {
		name        string
		args        args
		wantVisited [][]int
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{1, 0, 1, 1},
				},
				visited: [][]bool{
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
				x: 0,
				y: 0,
			},
			wantVisited: [][]int{
				{1, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 1, 1, 0},
				{0, 0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bfs(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)

			t.Logf("visitedPath: \n%s", prettyPrintGrid(tt.args.visited))
		})
	}
}

func Test_bfsII(t *testing.T) {
	type args struct {
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{1, 0, 1, 1},
				},
				visited: [][]bool{
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
				x: 0,
				y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bfsII(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)

			t.Logf("visitedPath: \n%s", prettyPrintGrid(tt.args.visited))
		})
	}
}

// go test -bench -cpu=2,4  -benchtime=5s -count=3 github.com/sword-fisher-fly/algorithm-and-structure/graph -v
// go test -bench -cpu=2,4  -benchtime=5s  github.com/sword-fisher-fly/algorithm-and-structure/graph -v
func BenchmarkBFS(b *testing.B) {
	type args struct {
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}

	tests := []struct {
		name        string
		args        args
		wantVisited [][]int
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{1, 0, 1, 1},
				},
				visited: [][]bool{
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
				x: 0,
				y: 0,
			},
		},
	}

	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			bfs(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)
			// b.Logf("visitedPath: \n%s", prettyPrintGrid(tt.args.visited))
		}
	}
}

func BenchmarkBFSII(b *testing.B) {
	type args struct {
		grid    [][]int
		visited [][]bool
		x       int
		y       int
	}

	tests := []struct {
		name        string
		args        args
		wantVisited [][]int
	}{
		{
			name: "4*4 grid from the top left corner to the bottom right corner",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{1, 0, 1, 1},
				},
				visited: [][]bool{
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
				x: 0,
				y: 0,
			},
		},
	}

	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			bfsII(tt.args.grid, tt.args.visited, tt.args.x, tt.args.y)
			// b.Logf("visitedPath: \n%s", prettyPrintGrid(tt.args.visited))
		}
	}
}
