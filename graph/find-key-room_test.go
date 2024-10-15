package graph

import (
	"testing"
)

func TestCanVisitAllRoomsByBFS(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can visit all rooms by bfs: 1",
			args: args{rooms: [][]int{{1}, {2}, {3}, {}}},
			want: true,
		},
		{
			name: "can visit all rooms by bfs: 2",
			args: args{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			want: false,
		},
		{
			name: "can visit all rooms by bfs: 3",
			args: args{rooms: [][]int{{5}, {}, {1, 3}, {5}, {}, {}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanVisitAllRoomsByBFS(tt.args.rooms); got != tt.want {
				t.Errorf("CanVisitAllRoomsByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanVisitAllRoomsByDFS(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can visit all rooms by dfs: 1",
			args: args{rooms: [][]int{{1}, {2}, {3}, {}}},
			want: true,
		},
		{
			name: "can visit all rooms by dfs: 2",
			args: args{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			want: false,
		},
		{
			name: "can visit all rooms by dfs: 3",
			args: args{rooms: [][]int{{5}, {}, {1, 3}, {5}, {}, {}}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanVisitAllRoomsByDFS(tt.args.rooms); got != tt.want {
				t.Errorf("CanVisitAllRoomsByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanVisitAllRoomsByDFSII(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can visit all rooms by dfs: 1",
			args: args{rooms: [][]int{{1}, {2}, {3}, {}}},
			want: true,
		},
		{
			name: "can visit all rooms by dfs: 2",
			args: args{rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			want: false,
		},
		{
			name: "can visit all rooms by dfs: 3",
			args: args{rooms: [][]int{{5}, {}, {1, 3}, {5}, {}, {}}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanVisitAllRoomsByDFSII(tt.args.rooms); got != tt.want {
				t.Errorf("CanVisitAllRoomsByDFSII() = %v, want %v", got, tt.want)
			}
		})
	}
}
