package dynamic

import (
	"testing"
)

func TestMinEditDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty word and real word",
			args: args{
				word1: "",
				word2: "hello",
			},
			want: 5,
		},
		{
			name: "eat and sea",
			args: args{
				word1: "eat",
				word2: "sea",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEditDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MinEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinEditDistanceII(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty word and real word",
			args: args{
				word1: "",
				word2: "hello",
			},
			want: 5,
		},
		{
			name: "eat and sea",
			args: args{
				word1: "eat",
				word2: "sea",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinEditDistanceII(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MinEditDistanceII() = %v, want %v", got, tt.want)
			}
		})
	}
}
