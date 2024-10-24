package trie

import (
	"testing"
)

func TestMinimumLengthEncodingII(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{words: []string{"time", "me", "bell"}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumLengthEncodingII(tt.args.words); got != tt.want {
				t.Errorf("MinimumLengthEncodingII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinimumLengthEncodingByTrie(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{words: []string{"time", "me", "bell"}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumLengthEncodingByTrie(tt.args.words); got != tt.want {
				t.Errorf("MinimumLengthEncodingByTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}
