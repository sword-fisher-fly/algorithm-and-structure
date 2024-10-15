package interview

import (
	"testing"
)

func TestLongestPalindromeByDP(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1: babad",
			args: args{
				s: "babad",
			},
			// want: "bab",
			want: "aba",
		},
		{
			name: "case 2: cbbd",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "case 3: a",
			args: args{
				s: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeByDP(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindromeByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeByExpand(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1: babad",
			args: args{
				s: "babad",
			},
			want: "bab",
			// want: "aba",
		},
		{
			name: "case 2: cbbd",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "case 3: a",
			args: args{
				s: "a",
			},
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindromeByExpand(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindromeByExpand() = %v, want %v", got, tt.want)
			}
		})
	}
}
