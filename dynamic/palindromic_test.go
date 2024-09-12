package dynamic

import (
	"testing"
)

func TestCountPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the same character",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
		{
			name: "different characters",
			args: args{
				s: "abac",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPalindromicSubsequence(tt.args.s); got != tt.want {
				t.Errorf("CountPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountPalindromicSubsequenceII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "the same character",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
		{
			name: "different characters",
			args: args{
				s: "abac",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPalindromicSubsequenceII(tt.args.s); got != tt.want {
				t.Errorf("CountPalindromicSubsequenceII() = %v, want %v", got, tt.want)
			}
		})
	}
}
