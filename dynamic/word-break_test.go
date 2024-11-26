package dynamic

import (
	"testing"
)

func TestWorkBreak(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "leetcode",
			args: args{
				s:     "leetcode",
				words: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "l love china",
			args: args{
				s:     "ilovechina",
				words: []string{"i", "love", "leetcode", "china"},
			},
			want: true,
		},
		{
			name: "applepenapple",
			args: args{
				s:     "applepenapple",
				words: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "catsandog",
			args: args{
				s:     "catsandog",
				words: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WorkBreak(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("WorkBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestWorkBreakII(t *testing.T) {
// 	type args struct {
// 		s     string
// 		words []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "leetcode",
// 			args: args{
// 				s:     "leetcode",
// 				words: []string{"leet", "code"},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "l love china",
// 			args: args{
// 				s:     "ilovechina",
// 				words: []string{"i", "love", "leetcode", "china"},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "applepenapple",
// 			args: args{
// 				s:     "applepenapple",
// 				words: []string{"apple", "pen"},
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "catsandog",
// 			args: args{
// 				s:     "catsandog",
// 				words: []string{"cats", "dog", "sand", "and", "cat"},
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := WorkBreakII(tt.args.s, tt.args.words); got != tt.want {
// 				t.Errorf("WorkBreakII() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestWordBreakByDFS(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "leetcode",
			args: args{
				s:     "leetcode",
				words: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "l love china",
			args: args{
				s:     "ilovechina",
				words: []string{"i", "love", "leetcode", "china"},
			},
			want: true,
		},
		{
			name: "applepenapple",
			args: args{
				s:     "applepenapple",
				words: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "catsandog",
			args: args{
				s:     "catsandog",
				words: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordBreakByDFS(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("WordBreakByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
