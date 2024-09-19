package dynamic

import "testing"

func TestMaxSubPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{s: "bbbab"},
			want: 4,
		},
		{
			name: "case2",
			args: args{s: "cbbd"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubPalindrome(tt.args.s); got != tt.want {
				t.Errorf("MaxSubPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
