package backtracking

import (
	"testing"
)

func TestSolveSudu(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sudu test",
			args: args{
				board: [][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveSudu(tt.args.board); got != tt.want {
				t.Errorf("SolveSudu() = %v, want %v", got, tt.want)
			}

			t.Logf("board: \n%+v", viewBoard(tt.args.board))
		})
	}
}

func TestIsValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				board: [][]byte{
					[]byte(".87654321"),
					[]byte("2........"),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				board: [][]byte{
					[]byte(".87654321"),
					[]byte("2.3......"),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("IsValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
