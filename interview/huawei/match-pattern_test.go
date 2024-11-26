package huawei

import (
	"testing"
)

func TestMatchPattern(t *testing.T) {
	type args struct {
		p    string
		pIdx int
		t    string
		tIdx int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{p: "te?t*.*", pIdx: 0, t: "txt12.xls", tIdx: 0},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				p:    "?*Bc*?",
				pIdx: 0,
				t:    "abcd",
				tIdx: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchPattern(tt.args.p, tt.args.pIdx, tt.args.t, tt.args.tIdx); got != tt.want {
				t.Errorf("MatchPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchPatternByDy(t *testing.T) {
	type args struct {
		p string
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{p: "te?t*.*", s: "txt12.xls"},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				p: "?*Bc*?",
				s: "abcd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchPatternByDy(tt.args.p, tt.args.s); got != tt.want {
				t.Errorf("MatchPatternByDy() = %v, want %v", got, tt.want)
			}
		})
	}
}
