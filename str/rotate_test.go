package str

import (
	"testing"
)

func TestRotateString(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n is less than the length of s",
			args: args{
				s: "abcdefg",
				n: 2,
			},
			want: "fgabcde",
		},
		{
			name: "n is greater than the length of s",
			args: args{
				s: "abcdefg",
				n: 8,
			},
			want: "gabcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateString(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("RotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotateStringII(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n is less than the length of s",
			args: args{
				s: "abcdefg",
				n: 2,
			},
			want: "fgabcde",
		},
		{
			name: "n is greater than the length of s",
			args: args{
				s: "abcdefg",
				n: 8,
			},
			want: "gabcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateStringII(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("RotateStringII() = %v, want %v", got, tt.want)
			}
		})
	}
}
