package backtracking

import "testing"

func TestWorkBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "leetcode",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "l love china",
			args: args{
				s:        "ilovechina",
				wordDict: []string{"i", "love", "leetcode", "china"},
			},
			want: true,
		},
		{
			name: "applepenapple",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "catsandog",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WorkBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("WorkBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
