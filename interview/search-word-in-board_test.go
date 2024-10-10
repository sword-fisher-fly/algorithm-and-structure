package interview

import (
	"testing"
)

func TestSearchWordInBoardByBFS(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "not found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABWORLD",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchWordInBoardByBFS(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("SearchWordInBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchWordInBoardByBacktracing(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "not found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABWORLD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchWordInBoardByBacktracing(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("SearchWordInBoardByBacktracing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchWordInBoardByDFS(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "not found a valid word path in board",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABWORLD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchWordInBoardByDFS(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("SearchWordInBoardByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
