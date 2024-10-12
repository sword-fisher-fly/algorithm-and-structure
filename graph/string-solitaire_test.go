package graph

import (
	"testing"
)

func TestStringSolitaire(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		words     []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exist two path that can reach the end word",
			args: args{
				beginWord: "abc",
				endWord:   "def",
				words:     []string{"efc", "dbc", "ebc", "dec", "dfc", "yhn"},
			},
			want: 4,
		},
		{
			name: "no exist path that can reach the end word",
			args: args{
				beginWord: "abc",
				endWord:   "zab",
				words:     []string{"efc", "dbc", "ebc", "dec", "dfc", "yhn"},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSolitaire(tt.args.beginWord, tt.args.endWord, tt.args.words); got != tt.want {
				t.Errorf("StringSolitaire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSolitaireByDFS(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		words     []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exist two path that can reach the end word",
			args: args{
				beginWord: "abc",
				endWord:   "def",
				words:     []string{"efc", "dbc", "ebc", "dec", "dfc", "yhn"},
			},
			want: 4,
		},
		// {
		// 	name: "no exist path that can reach the end word",
		// 	args: args{
		// 		beginWord: "abc",
		// 		endWord:   "zab",
		// 		words:     []string{"efc", "dbc", "ebc", "dec", "dfc", "yhn"},
		// 	},
		// 	want: 0,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSolitaireByDFS(tt.args.beginWord, tt.args.endWord, tt.args.words); got != tt.want {
				t.Errorf("StringSolitaireByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
